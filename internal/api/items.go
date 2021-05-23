package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"

	openapi "github.com/atye/gosrsbox/internal/openapi/api"
	"github.com/atye/gosrsbox/osrsbox"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
	"golang.org/x/sync/errgroup"
)

func (c *client) GetItemsByID(ctx context.Context, ids ...string) ([]osrsbox.Item, error) {
	if len(ids) == 0 {
		return nil, errors.New("no ids provided")
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *client) GetItemsByName(ctx context.Context, names ...string) ([]osrsbox.Item, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *client) GetItemSet(ctx context.Context, set sets.SetName) ([]osrsbox.Item, error) {
	if set == nil || len(set) == 0 {
		return nil, errors.New("no set provided")
	}
	return c.GetItemsByName(ctx, set...)
}

func (c *client) GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]osrsbox.Item, error) {
	if slot == "" || len(slot) == 0 {
		return nil, errors.New("no set provided")
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "equipable_by_player": true, "equipment.slot": %s, "duplicate": false }`, slot))
}

func (c *client) GetItemsByQuery(ctx context.Context, query string) ([]osrsbox.Item, error) {
	resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.ItemApi.Getitems(ctx).Where(query))
	if err != nil {
		return nil, err
	}

	switch inline := resp.(type) {
	case openapi.InlineResponse200:
		pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
		items := make([]osrsbox.Item, *inline.Meta.Total)

		respItems, err := openItemsToItems(inline.GetItems())
		if err != nil {
			return nil, err
		}

		_ = copy(items, respItems)

		if pages > 1 {
			var eg errgroup.Group
			for page := 2; page <= pages; page++ {
				page := page
				eg.Go(func() error {
					resp, err := c.retryOpenAPIRequest(ctx, c.openAPIClient.ItemApi.Getitems(ctx).Where(query).Page(int32(page)), errTooManyOpenFiles)
					if err != nil {
						return err
					}

					switch inline := resp.(type) {
					case openapi.InlineResponse200:
						respItems, err := openItemsToItems(inline.GetItems())
						if err != nil {
							return err
						}

						for i, item := range respItems {
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

func openItemsToItems(openItems []openapi.Item) ([]osrsbox.Item, error) {
	b, err := json.Marshal(openItems)
	if err != nil {
		return nil, err
	}

	var items []osrsbox.Item
	err = json.Unmarshal(b, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
