package client

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/atye/gosrsbox/internal/common"
	"github.com/atye/gosrsbox/models"
	"golang.org/x/sync/errgroup"
)

func (c *APIClient) GetPrayersByID(ctx context.Context, ids ...string) ([]models.Prayer, error) {
	ctx, span := c.createSpan(ctx, "get_prayers_by_id")
	defer span.End()

	if len(ids) == 0 {
		return nil, errors.New("no ids provided")
	}
	return c.GetPrayersByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }}`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *APIClient) GetPrayersByName(ctx context.Context, names ...string) ([]models.Prayer, error) {
	ctx, span := c.createSpan(ctx, "get_prayers_by_name")
	defer span.End()

	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}

	query := fmt.Sprintf(`{ "name": { "$in": [%s] } }`, strings.Join(quoteStrings(names...), ", "))
	return c.GetPrayersByQuery(ctx, query)
}

func (c *APIClient) GetPrayersByQuery(ctx context.Context, query string) ([]models.Prayer, error) {
	ctx, span := c.createSpan(ctx, "get_prayers_by_query")
	defer span.End()

	inline, err := c.doPrayersRequest(ctx, common.Params{Where: query})
	if err != nil {
		return nil, err
	}

	total := inline.GetTotal()
	maxResults := inline.GetMaxResults()

	pages := int(math.Ceil(float64(total) / float64(maxResults)))
	prayers := make([]models.Prayer, total)

	_ = copy(prayers, inline.GetPrayers())

	if pages > 1 {
		var eg errgroup.Group
		for page := 2; page <= pages; page++ {
			page := page
			eg.Go(func() error {
				inline, err := c.doPrayersRequest(ctx, common.Params{Where: query, Page: page})
				if err != nil {
					return err
				}

				for i, prayer := range inline.GetPrayers() {
					prayers[maxResults*(page-1)+i] = prayer
				}

				return nil
			})
		}
		err := eg.Wait()
		if err != nil {
			return nil, err
		}
	}
	return prayers, nil
}
