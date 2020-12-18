package api

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

type APIClient struct {
	address string
	client  *http.Client
	mu      sync.Mutex
}

const (
	api              = "https://api.osrsbox.com"
	itemsEndpoint    = "items"
	monstersEndpoint = "monsters"
	prayersEndpoint  = "prayers"
)

func NewAPIClient() *APIClient {
	return &APIClient{
		address: api,
		client:  http.DefaultClient,
		mu:      sync.Mutex{},
	}
}

func (c *APIClient) doRequest(ctx context.Context, url string, v interface{}) error {
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
