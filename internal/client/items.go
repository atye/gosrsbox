package client

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/atye/gosrsbox/models"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
	"golang.org/x/sync/errgroup"
)

func (c *apiClient) GetItemsByID(ctx context.Context, ids ...string) ([]models.Item, error) {
	if len(ids) == 0 {
		return nil, errNoIDs
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *apiClient) GetItemsByName(ctx context.Context, names ...string) ([]models.Item, error) {
	if len(names) == 0 {
		return nil, errNoNames
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *apiClient) GetItemSet(ctx context.Context, set sets.SetName) ([]models.Item, error) {
	if set == nil || len(set) == 0 {
		return nil, errNoSet
	}
	return c.GetItemsByName(ctx, set...)
}

func (c *apiClient) GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]models.Item, error) {
	if slot == "" {
		return nil, errNoSlot
	}
	return c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "equipable_by_player": true, "equipment.slot": %s, "duplicate": false }`, slot))
}

func (c *apiClient) GetItemsByQuery(ctx context.Context, query string) ([]models.Item, error) {
	inline, err := c.doItemsRequest(ctx, c.openAPIClient.ItemApi.Getitems(ctx).Where(query))
	if err != nil {
		return nil, err
	}

	pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
	items := make([]models.Item, *inline.Meta.Total)

	/*var tmpItems []models.Item
	err = convert(inline.GetItems(), &tmpItems)
	if err != nil {
		return nil, err
	}*/

	_ = copy(items, inline.GetItems())

	if pages > 1 {
		var eg errgroup.Group
		for page := 2; page <= pages; page++ {
			page := page
			eg.Go(func() error {
				inline, err := c.doItemsRequest(ctx, c.openAPIClient.ItemApi.Getitems(ctx).Where(query).Page(int32(page)))
				if err != nil {
					return err
				}

				/*var tmpItems []models.Item
				err = convert(inline.GetItems(), &tmpItems)
				if err != nil {
					return err
				}*/

				for i, item := range inline.GetItems() {
					items[int(*inline.Meta.MaxResults)*(page-1)+i] = item
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