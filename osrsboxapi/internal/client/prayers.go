package client

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"

	openapi "github.com/atye/gosrsbox/osrsboxapi/openapi/client"
	"golang.org/x/sync/errgroup"
)

func (c *client) GetPrayersByID(ctx context.Context, ids ...string) ([]openapi.Prayer, error) {
	if len(ids) == 0 {
		return nil, errors.New("no ids provided")
	}
	return c.GetPrayersByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }}`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *client) GetPrayersByName(ctx context.Context, names ...string) ([]openapi.Prayer, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}

	query := fmt.Sprintf(`{ "name": { "$in": [%s] } }`, strings.Join(quoteStrings(names...), ", "))
	return c.GetPrayersByQuery(ctx, query)
}

func (c *client) GetPrayersByQuery(ctx context.Context, query string) ([]openapi.Prayer, error) {
	resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.PrayerApi.Getprayers(ctx).Where(query))
	if err != nil {
		return nil, err
	}
	switch inline := resp.(type) {
	case openapi.InlineResponse2004:
		pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
		prayers := make([]openapi.Prayer, *inline.Meta.Total)
		for i, prayer := range inline.GetItems() {
			prayers[i] = prayer
		}
		if pages > 1 {
			var eg errgroup.Group
			for page := 2; page <= pages; page++ {
				page := page
				eg.Go(func() error {
					resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.PrayerApi.Getprayers(ctx).Where(query).Page(int32(page)))
					if err != nil {
						return err
					}
					if inline, ok := resp.(openapi.InlineResponse2004); ok {
						for i, prayer := range inline.GetItems() {
							// check if something already exists?
							prayers[int(*inline.Meta.MaxResults)*(page-1)+i] = prayer
						}
					} else {
						return fmt.Errorf("unexpected inline item type type %T", inline)
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
	default:
		return nil, fmt.Errorf("unexpected response type %T", inline)
	}
}
