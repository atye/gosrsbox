package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/atye/gosrsbox/openapi/api"
	"golang.org/x/sync/semaphore"
)

type client struct {
	docsAddress   string
	openAPIClient *api.APIClient
	sem           *semaphore.Weighted
}

type ItemRequestExecutor interface {
	Execute() (api.InlineResponse200, *http.Response, api.GenericOpenAPIError)
}

type MonsterRequestExecutor interface {
	Execute() (api.InlineResponse2003, *http.Response, api.GenericOpenAPIError)
}

type PrayerRequestExecutor interface {
	Execute() (api.InlineResponse2004, *http.Response, api.GenericOpenAPIError)
}

const (
	jsonDocuments = "https://www.osrsbox.com/osrsbox-db/"
)

var (
	errNoIDs   = errors.New("no ids provided")
	errNoNames = errors.New("no names provided")
	errNoSet   = errors.New("no set provided")
	errNoSlot  = errors.New("no slot provided")
)

func NewAPI(conf *api.Configuration) *client {
	return &client{
		docsAddress:   jsonDocuments,
		openAPIClient: api.NewAPIClient(conf),
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
	case ItemRequestExecutor:
		inline, resp, openAPIErr := r.Execute()
		err := checkError(openAPIErr)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return inline, nil
	case MonsterRequestExecutor:
		inline, resp, openAPIErr := r.Execute()
		err := checkError(openAPIErr)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return inline, nil
	case PrayerRequestExecutor:
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

func (c *client) doDocumentRequest(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		return nil, fmt.Errorf("code: %d, message: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return resp, nil
}

func checkError(openAPIErr api.GenericOpenAPIError) error {
	if openAPIErr.Error() == "" {
		return nil
	}

	var apiErr api.Error
	err := json.Unmarshal(openAPIErr.Body(), &apiErr)
	if err != nil {
		return openAPIErr
	}

	if apiErr.Error.GetCode() == 0 && apiErr.Error.GetMessage() == "" {
		return openAPIErr
	}
	return fmt.Errorf("code %d, message: %s", apiErr.Error.GetCode(), apiErr.Error.GetMessage())
}

func quoteStrings(elements ...string) []string {
	quotedStrings := make([]string, len(elements))
	for i, e := range elements {
		quotedStrings[i] = fmt.Sprintf(`"%s"`, e)
	}
	return quotedStrings
}
