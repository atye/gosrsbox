package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/atye/gosrsbox/internal/api"
	"golang.org/x/sync/semaphore"
)

type apiClient struct {
	docsAddress   string
	openAPIClient *api.APIClient
	sem           *semaphore.Weighted
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

func NewAPI(conf *api.Configuration) *apiClient {
	return &apiClient{
		docsAddress:   jsonDocuments,
		openAPIClient: api.NewAPIClient(conf),
		sem:           semaphore.NewWeighted(int64(10)),
	}
}

func (c *apiClient) doItemsRequest(ctx context.Context, req api.ApiGetitemsRequest) (api.InlineResponse200, error) {
	err := c.sem.Acquire(ctx, 1)
	if err != nil {
		return api.InlineResponse200{}, err
	}
	defer c.sem.Release(1)

	inline, resp, err := req.Execute()
	err = checkError(err)
	if err != nil {
		return api.InlineResponse200{}, err
	}
	defer resp.Body.Close()

	return inline, nil
}

func (c *apiClient) doMonstersRequest(ctx context.Context, req api.ApiGetmonstersRequest) (api.InlineResponse2003, error) {
	err := c.sem.Acquire(ctx, 1)
	if err != nil {
		return api.InlineResponse2003{}, err
	}
	defer c.sem.Release(1)

	inline, resp, err := req.Execute()
	err = checkError(err)
	if err != nil {
		return api.InlineResponse2003{}, err
	}
	defer resp.Body.Close()

	return inline, nil
}

func (c *apiClient) doPrayersRequest(ctx context.Context, req api.ApiGetprayersRequest) (api.InlineResponse2004, error) {
	err := c.sem.Acquire(ctx, 1)
	if err != nil {
		return api.InlineResponse2004{}, err
	}
	defer c.sem.Release(1)

	inline, resp, err := req.Execute()
	err = checkError(err)
	if err != nil {
		return api.InlineResponse2004{}, err
	}
	defer resp.Body.Close()

	return inline, nil
}

func (c *apiClient) doDocumentRequest(ctx context.Context, url string) (*http.Response, error) {
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

func checkError(executeErr error) error {
	if executeErr == nil {
		return nil
	}

	var genericErr api.GenericOpenAPIError
	if !errors.As(executeErr, &genericErr) {
		return executeErr
	}

	var apiErr api.Error
	err := json.Unmarshal(genericErr.Body(), &apiErr)
	if err != nil {
		return err
	}

	if apiErr.Error.GetCode() == 0 && apiErr.Error.GetMessage() == "" {
		return executeErr
	}
	return fmt.Errorf("code %d, message: %s", apiErr.Error.GetCode(), apiErr.Error.GetMessage())
}

func quoteStrings(elements ...string) []string {
	quotedStrings := make([]string, len(elements))
	for i, e := range elements {
		quotedStrings[i] = fmt.Sprintf("%q", e)
	}
	return quotedStrings
}
