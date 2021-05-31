package openapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/atye/gosrsbox/internal/common"
	"github.com/atye/gosrsbox/internal/openapi/api"
	"github.com/atye/gosrsbox/models"
)

type client struct {
	openAPIClient *api.APIClient
}

func NewClient(userAgent string, scheme string, url string) *client {
	conf := &api.Configuration{
		Scheme:     scheme,
		HTTPClient: http.DefaultClient,
		UserAgent:  userAgent,
		Servers: []api.ServerConfiguration{
			{
				URL: url,
			},
		},
	}

	return &client{
		openAPIClient: api.NewAPIClient(conf),
	}
}

func (c *client) ExecuteItemsRequest(ctx context.Context, p common.Params) (common.ItemsResponse, error) {
	r := c.openAPIClient.ItemApi.Getitems(ctx).Where(p.Where)
	if p.Page != 0 {
		r = r.Page(int32(p.Page))
	}

	inline, resp, err := r.Execute()
	err = checkError(err)
	if err != nil {
		return itemsResponse{}, err
	}
	defer resp.Body.Close()

	b, err := json.Marshal(inline.GetItems())
	if err != nil {
		return itemsResponse{}, nil
	}

	var items []models.Item
	err = json.Unmarshal(b, &items)
	if err != nil {
		return itemsResponse{}, nil
	}

	return itemsResponse{
		total:      int(*inline.Meta.Total),
		maxResults: int(*inline.Meta.MaxResults),
		items:      items,
	}, nil
}

func (c *client) ExecuteMonstersRequest(ctx context.Context, p common.Params) (common.MonstersResponse, error) {
	r := c.openAPIClient.MonsterApi.Getmonsters(ctx).Where(p.Where)
	if p.Page != 0 {
		r = r.Page(int32(p.Page))
	}

	inline, resp, err := r.Execute()
	err = checkError(err)
	if err != nil {
		return monstersResponse{}, err
	}
	defer resp.Body.Close()

	b, err := json.Marshal(inline.GetItems())
	if err != nil {
		return monstersResponse{}, nil
	}

	var monsters []models.Monster
	err = json.Unmarshal(b, &monsters)
	if err != nil {
		return monstersResponse{}, nil
	}

	return monstersResponse{
		total:      int(*inline.Meta.Total),
		maxResults: int(*inline.Meta.MaxResults),
		monsters:   monsters,
	}, nil
}

func (c *client) ExecutePrayersRequest(ctx context.Context, p common.Params) (common.PrayersResponse, error) {
	r := c.openAPIClient.PrayerApi.Getprayers(ctx).Where(p.Where)
	if p.Page != 0 {
		r = r.Page(int32(p.Page))
	}

	inline, resp, err := r.Execute()
	err = checkError(err)
	if err != nil {
		return prayersResponse{}, err
	}
	defer resp.Body.Close()

	b, err := json.Marshal(inline.GetItems())
	if err != nil {
		return prayersResponse{}, nil
	}

	var prayers []models.Prayer
	err = json.Unmarshal(b, &prayers)
	if err != nil {
		return prayersResponse{}, nil
	}

	return prayersResponse{
		total:      int(*inline.Meta.Total),
		maxResults: int(*inline.Meta.MaxResults),
		prayers:    prayers,
	}, nil
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

type itemsResponse struct {
	total      int
	maxResults int
	items      []models.Item
}

func (r itemsResponse) GetTotal() int {
	return r.total
}

func (r itemsResponse) GetMaxResults() int {
	return r.maxResults
}

func (r itemsResponse) GetItems() []models.Item {
	return r.items
}

type monstersResponse struct {
	total      int
	maxResults int
	monsters   []models.Monster
}

func (r monstersResponse) GetTotal() int {
	return r.total
}

func (r monstersResponse) GetMaxResults() int {
	return r.maxResults
}

func (r monstersResponse) GetMonsters() []models.Monster {
	return r.monsters
}

type prayersResponse struct {
	total      int
	maxResults int
	prayers    []models.Prayer
}

func (r prayersResponse) GetTotal() int {
	return r.total
}

func (r prayersResponse) GetMaxResults() int {
	return r.maxResults
}

func (r prayersResponse) GetPrayers() []models.Prayer {
	return r.prayers
}
