package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/sync/semaphore"
)

type client struct {
	address string
	client  *http.Client
	sem     *semaphore.Weighted
	mu      sync.Mutex
}

const (
	api              = "https://api.osrsbox.com"
	itemsEndpoint    = "items"
	monstersEndpoint = "monsters"
	prayersEndpoint  = "prayers"
)

func NewAPI(httpClient *http.Client) *client {
	return &client{
		address: api,
		client:  httpClient,
		sem:     semaphore.NewWeighted(int64(10)),
		mu:      sync.Mutex{},
	}
}

func (c *client) doRequest(ctx context.Context, url string, v interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	err = c.sem.Acquire(ctx, 1)
	if err != nil {
		return err
	}
	defer c.sem.Release(1)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
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
