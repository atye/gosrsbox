package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	openapi "github.com/atye/gosrsbox/osrsboxapi/openapi/client"
)

type client struct {
	docsAddress   string
	httpClient    *http.Client
	openAPIClient *openapi.APIClient
	mu            sync.Mutex
}

const (
	docs = "https://www.osrsbox.com/osrsbox-db"
)

func NewAPI(conf *openapi.Configuration) *client {
	return &client{
		docsAddress:   docs,
		httpClient:    conf.HTTPClient,
		openAPIClient: openapi.NewAPIClient(conf),
	}
}

func (c *client) doOpenAPIRequest(ctx context.Context, req interface{}) (interface{}, error) {
	switch r := req.(type) {
	case openapi.ApiGetitemsRequest:
		inline, resp, openAPIErr := r.Execute()
		defer resp.Body.Close()
		err := checkOpenAPIErr(resp, openAPIErr)
		if err != nil {
			return nil, err
		}
		return inline, nil
	case openapi.ApiGetmonstersRequest:
		inline, resp, openAPIErr := r.Execute()
		defer resp.Body.Close()
		err := checkOpenAPIErr(resp, openAPIErr)
		if err != nil {
			return nil, err
		}
		return inline, nil
	case openapi.ApiGetprayersRequest:
		inline, resp, openAPIErr := r.Execute()
		defer resp.Body.Close()
		err := checkOpenAPIErr(resp, openAPIErr)
		if err != nil {
			return nil, err
		}
		return inline, nil
	default:
		return nil, fmt.Errorf("request type %T not supported", r)
	}
}

func checkOpenAPIErr(resp *http.Response, err openapi.GenericOpenAPIError) error {
	if err.Body() != nil {
		var apiErr openapi.Error
		err := json.NewDecoder(resp.Body).Decode(&apiErr)
		if err != nil {
			return err
		}
		return fmt.Errorf("code: %d, message: %s", *apiErr.GetError().Code, *apiErr.GetError().Message)
	}
	return nil
}

func (c *client) doJSONDocsRequest(ctx context.Context, url string, v interface{}) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		if resp != nil {
			return resp.StatusCode, err
		}
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("expected status 200/OK, got %d", resp.StatusCode)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return resp.StatusCode, err
	}
	return resp.StatusCode, nil
}

func quoteStrings(elements ...string) []string {
	quotedStrings := make([]string, len(elements))
	for i, e := range elements {
		quotedStrings[i] = fmt.Sprintf(`"%s"`, e)
	}
	return quotedStrings
}
