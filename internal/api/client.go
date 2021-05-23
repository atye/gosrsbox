package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	openapi "github.com/atye/gosrsbox/internal/openapi/api"
	"golang.org/x/sync/semaphore"
)

type client struct {
	docsAddress   string
	openAPIClient *openapi.APIClient
	sem           *semaphore.Weighted
}

type apiError struct {
	Status string `json:"_status"`
	Err    struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"_error"`
}

func (e apiError) Error() string {
	return fmt.Sprintf("code %d, message: %s", e.Err.Code, e.Err.Message)
}

var (
	errTooManyOpenFiles = "too many open files"
)

const (
	jsonDocuments = "https://www.osrsbox.com/osrsbox-db/"
)

func NewAPI(conf *openapi.Configuration) *client {
	return &client{
		docsAddress:   jsonDocuments,
		openAPIClient: openapi.NewAPIClient(conf),
		sem:           semaphore.NewWeighted(int64(10)),
	}
}

func (c *client) doOpenAPIRequest(ctx context.Context, req interface{}) (interface{}, error) {
	err := c.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer c.sem.Release(1)

	switch r := req.(type) {
	case openapi.ApiGetitemsRequest:
		inline, resp, openAPIErr := r.Execute()
		err := checkError(openAPIErr)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return inline, nil
	case openapi.ApiGetmonstersRequest:
		inline, resp, openAPIErr := r.Execute()
		err := checkError(openAPIErr)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return inline, nil
	case openapi.ApiGetprayersRequest:
		inline, resp, openAPIErr := r.Execute()
		err := checkError(openAPIErr)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return inline, nil
	default:
		return nil, fmt.Errorf("request type %T not supported", r)
	}
}

func (c *client) retryOpenAPIRequest(ctx context.Context, req interface{}, errMessages ...string) (interface{}, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	resp, err := c.doOpenAPIRequest(ctx, req)
	if err != nil {
		for _, m := range errMessages {
			if strings.Contains(err.Error(), m) {
				time.Sleep(1 * time.Second)
				return c.retryOpenAPIRequest(ctx, req)
			}
		}
		return nil, err
	}
	return resp, nil
}

func (c *client) doDocumentRequest(ctx context.Context, url string, v interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("reading response body: %w", err)
		}
		return fmt.Errorf("code: %d, message: %s", resp.StatusCode, string(body))
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, v)
	if err != nil {
		return err
	}
	return nil
}

func quoteStrings(elements ...string) []string {
	quotedStrings := make([]string, len(elements))
	for i, e := range elements {
		quotedStrings[i] = fmt.Sprintf(`"%s"`, e)
	}
	return quotedStrings
}

func checkError(openAPIErr openapi.GenericOpenAPIError) error {
	if openAPIErr.Error() != "" {
		var apiErr apiError
		err := json.Unmarshal(openAPIErr.Body(), &apiErr)
		if err != nil {
			return openAPIErr
		}

		if apiErr.Err.Code == 0 && apiErr.Err.Message == "" {
			return openAPIErr
		}
		return apiErr
	}
	return nil
}
