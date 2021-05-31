package client

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/atye/gosrsbox/internal/common"
	"github.com/atye/gosrsbox/models"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
	"golang.org/x/sync/errgroup"
)

func (c *APIClient) GetItemsByID(ctx context.Context, ids ...string) ([]models.Item, error) {
	ctx, span := c.createSpan(ctx, "get_items_by_id")
	defer span.End()

	var err error
	defer func() {
		if err != nil {
			setSpanErrorStatus(span, err)
		}
	}()

	if len(ids) == 0 {
		err = errNoIDs
		return nil, err
	}

	items, err := c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (c *APIClient) GetItemsByName(ctx context.Context, names ...string) ([]models.Item, error) {
	ctx, span := c.createSpan(ctx, "get_items_by_name")
	defer span.End()

	var err error
	defer func() {
		if err != nil {
			setSpanErrorStatus(span, err)
		}
	}()

	if len(names) == 0 {
		err = errNoNames
		return nil, err
	}

	items, err := c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (c *APIClient) GetItemSet(ctx context.Context, set sets.SetName) ([]models.Item, error) {
	ctx, span := c.createSpan(ctx, "get_item_set")
	defer span.End()

	var err error
	defer func() {
		if err != nil {
			setSpanErrorStatus(span, err)
		}
	}()

	if set == nil || len(set) == 0 {
		err = errNoSet
		return nil, err
	}
	items, err := c.GetItemsByName(ctx, set...)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (c *APIClient) GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]models.Item, error) {
	ctx, span := c.createSpan(ctx, "get_items_by_slot")
	defer span.End()

	var err error
	defer func() {
		if err != nil {
			setSpanErrorStatus(span, err)
		}
	}()

	if slot == "" {
		err = errNoSlot
		return nil, err
	}
	items, err := c.GetItemsByQuery(ctx, fmt.Sprintf(`{ "equipable_by_player": true, "equipment.slot": %s, "duplicate": false }`, slot))
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (c *APIClient) GetItemsByQuery(ctx context.Context, query string) ([]models.Item, error) {
	ctx, span := c.createSpan(ctx, "get_items_by_query")
	defer span.End()

	var err error
	defer func() {
		if err != nil {
			setSpanErrorStatus(span, err)
		}
	}()

	inline, err := c.doItemsRequest(ctx, common.Params{Where: query})
	if err != nil {
		return nil, err
	}

	total := inline.GetTotal()
	maxResults := inline.GetMaxResults()

	pages := int(math.Ceil(float64(total) / float64(maxResults)))
	items := make([]models.Item, total)

	_ = copy(items, inline.GetItems())

	if pages > 1 {
		var eg errgroup.Group
		for page := 2; page <= pages; page++ {
			page := page
			eg.Go(func() error {
				inline, err := c.doItemsRequest(ctx, common.Params{Where: query, Page: page})
				if err != nil {
					return err
				}

				for i, item := range inline.GetItems() {
					items[maxResults*(page-1)+i] = item
				}

				return nil
			})
		}
		err = eg.Wait()
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}
