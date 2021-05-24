package client

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/atye/gosrsbox/openapi/api"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
	"golang.org/x/sync/errgroup"
)

func (c *client) GetItemsByID(ctx context.Context, ids ...string) ([]api.Item, error) {
	if len(ids) == 0 {
		return nil, errNoIDs
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *client) GetItemsByName(ctx context.Context, names ...string) ([]api.Item, error) {
	if len(names) == 0 {
		return nil, errNoNames
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *client) GetItemSet(ctx context.Context, set sets.SetName) ([]api.Item, error) {
	if set == nil || len(set) == 0 {
		return nil, errNoSet
	}
	return c.GetItemsByName(ctx, set...)
}

func (c *client) GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]api.Item, error) {
	if slot == "" {
		return nil, errNoSlot
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "equipable_by_player": true, "equipment.slot": %s, "duplicate": false }`, slot))
}

func (c *client) GetItemsByQuery(ctx context.Context, query string) ([]api.Item, error) {
	resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.ItemApi.Getitems(ctx).Where(query))
	if err != nil {
		return nil, err
	}

	switch inline := resp.(type) {
	case api.InlineResponse200:
		pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
		items := make([]api.Item, *inline.Meta.Total)

		_ = copy(items, inline.GetItems())

		if pages > 1 {
			var eg errgroup.Group
			for page := 2; page <= pages; page++ {
				page := page
				eg.Go(func() error {
					resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.ItemApi.Getitems(ctx).Where(query).Page(int32(page)))
					if err != nil {
						return err
					}

					switch inline := resp.(type) {
					case api.InlineResponse200:
						for i, item := range inline.GetItems() {
							items[int(*inline.Meta.MaxResults)*(page-1)+i] = item
						}
					default:
						return fmt.Errorf("unexpected response type %T", inline)
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
