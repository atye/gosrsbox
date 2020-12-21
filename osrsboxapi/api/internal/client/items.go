package client

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/url"
	"strings"

	"github.com/atye/gosrsbox/osrsboxapi"
	"github.com/atye/gosrsbox/osrsboxapi/sets"
	"github.com/atye/gosrsbox/osrsboxapi/slots"
	"golang.org/x/sync/errgroup"
)

func (c *client) GetItemsByName(ctx context.Context, names ...string) ([]osrsboxapi.Item, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *client) GetItemSet(ctx context.Context, set sets.SetName) ([]osrsboxapi.Item, error) {
	if set == nil || len(set) == 0 {
		return nil, errors.New("no set provided")
	}
	return c.GetItemsByName(ctx, set...)
}

func (c *client) GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]osrsboxapi.Item, error) {
	if slot == "" || len(slot) == 0 {
		return nil, errors.New("no set provided")
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "equipable_by_player": true, "equipment.slot": %s, "duplicate": false }`, slot))
}

func (c *client) GetItemsByQuery(ctx context.Context, query string) ([]osrsboxapi.Item, error) {
	apiURL := fmt.Sprintf("%s/%s?where=%s", c.apiAddress, itemsEndpoint, url.QueryEscape(query))

	var itemsResp itemsResponse
	_, err := c.doAPIRequest(ctx, apiURL, &itemsResp)
	if itemsResp.Error != nil {
		return nil, itemsResp.Error
	}
	if err != nil {
		return nil, err
	}

	items := make([]osrsboxapi.Item, itemsResp.Meta.Total)
	for i, item := range itemsResp.Items {
		items[i] = item
	}

	var pages int
	if itemsResp.Meta.MaxResults != 0 {
		pages = int(math.Ceil(float64(itemsResp.Meta.Total) / float64(itemsResp.Meta.MaxResults)))
	}

	if pages > 1 {
		var eg errgroup.Group
		for page := 2; page <= pages; page++ {
			page := page
			eg.Go(func() error {
				var temp itemsResponse
				_, err := c.doAPIRequest(ctx, fmt.Sprintf("%s%s", apiURL, fmt.Sprintf("&page=%d", page)), &temp)
				if temp.Error != nil {
					return temp.Error
				}
				if err != nil {
					return err
				}
				for i, item := range temp.Items {
					// check if something already exists?
					items[temp.Meta.MaxResults*(page-1)+i] = item
				}
				return nil
			})
		}
		err := eg.Wait()
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}
