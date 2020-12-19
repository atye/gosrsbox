package external

import (
	"context"
	"encoding/json"
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

func NewAPI() *client {
	return &client{
		address: api,
		client:  http.DefaultClient,
		sem:     semaphore.NewWeighted(int64(10)),
		mu:      sync.Mutex{},
	}
}

func (c *client) doRequest(ctx context.Context, url string, v interface{}) error {
	err := c.sem.Acquire(ctx, 1)
	if err != nil {
		return err
	}
	defer c.sem.Release(1)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
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
