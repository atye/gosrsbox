package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"golang.org/x/sync/semaphore"
)

type client struct {
	apiAddress  string
	docsAddress string
	client      *http.Client
	apiSemaphor *semaphore.Weighted
	mu          sync.Mutex
}

const (
	docs             = "https://www.osrsbox.com/osrsbox-db"
	api              = "https://api.osrsbox.com"
	itemsEndpoint    = "items"
	monstersEndpoint = "monsters"
	prayersEndpoint  = "prayers"
)

func NewAPI(httpClient *http.Client) *client {
	return &client{
		apiAddress:  api,
		docsAddress: docs,
		client:      httpClient,
		apiSemaphor: semaphore.NewWeighted(int64(10)),
		mu:          sync.Mutex{},
	}
}

func (c *client) doAPIRequest(ctx context.Context, url string, v interface{}) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	err = c.apiSemaphor.Acquire(ctx, 1)
	if err != nil {
		return 0, err
	}
	defer c.apiSemaphor.Release(1)
	resp, err := c.client.Do(req)
	if err != nil {
		if resp != nil {
			return resp.StatusCode, err
		}
		return 0, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return resp.StatusCode, err
	}
	return resp.StatusCode, nil
}

func (c *client) doJSONDocsRequest(ctx context.Context, url string, v interface{}) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		if resp != nil {
			return resp.StatusCode, err
		}
		return 0, err
	}
	defer resp.Body.Close()
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
