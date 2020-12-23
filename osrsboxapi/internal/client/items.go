package client

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"

	openapi "github.com/atye/gosrsbox/osrsboxapi/openapi/client"
	"github.com/atye/gosrsbox/osrsboxapi/sets"
	"github.com/atye/gosrsbox/osrsboxapi/slots"
	"golang.org/x/sync/errgroup"
)

func (c *client) GetItemsByID(ctx context.Context, ids ...string) ([]openapi.Item, error) {
	if len(ids) == 0 {
		return nil, errors.New("no ids provided")
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *client) GetItemsByName(ctx context.Context, names ...string) ([]openapi.Item, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *client) GetItemSet(ctx context.Context, set sets.SetName) ([]openapi.Item, error) {
	if set == nil || len(set) == 0 {
		return nil, errors.New("no set provided")
	}
	return c.GetItemsByName(ctx, set...)
}

func (c *client) GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]openapi.Item, error) {
	if slot == "" || len(slot) == 0 {
		return nil, errors.New("no set provided")
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "equipable_by_player": true, "equipment.slot": %s, "duplicate": false }`, slot))
}

func (c *client) GetItemsByQuery(ctx context.Context, query string) ([]openapi.Item, error) {
	resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.ItemApi.Getitems(ctx).Where(query))
	if err != nil {
		return nil, err
	}
	switch inline := resp.(type) {
	case openapi.InlineResponse200:
		pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
		items := make([]openapi.Item, *inline.Meta.Total)
		for i, item := range inline.GetItems() {
			items[i] = item
		}
		if pages > 1 {
			var eg errgroup.Group
			for page := 2; page <= pages; page++ {
				page := page
				eg.Go(func() error {
					resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.ItemApi.Getitems(ctx).Where(query).Page(int32(page)))
					if err != nil {
						return err
					}
					if inline, ok := resp.(openapi.InlineResponse200); ok {
						for i, item := range inline.GetItems() {
							// check if something already exists?
							items[int(*inline.Meta.MaxResults)*(page-1)+i] = item
						}
					} else {
						return fmt.Errorf("unexpected inline item type %T", inline)
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
	default:
		return nil, fmt.Errorf("unexpected response type %T", inline)
	}
}
