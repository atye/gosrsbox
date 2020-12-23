package client

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/atye/gosrsbox/osrsboxapi/api/internal/client/openapi"
	"golang.org/x/sync/errgroup"
)

func (c *client) GetPrayersByName(ctx context.Context, names ...string) ([]openapi.Prayer, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	query := fmt.Sprintf(`{ "name": { "$in": [%s] } }`, strings.Join(quoteStrings(names...), ", "))
	return c.GetPrayersByQuery(ctx, query)
}

func (c *client) GetPrayersByQuery(ctx context.Context, query string) ([]openapi.Prayer, error) {
	resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.PrayerApi.Getprayers(ctx).Where(query))
	if err != nil {
		return nil, err
	}
	var pages int
	var inline openapi.InlineResponse2004
	if inline, ok := resp.(openapi.InlineResponse2004); ok {
		pages = int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
	} else {
		return nil, fmt.Errorf("unexpected type %T", inline)
	}
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
					for i, prayer := range *inline.Items {
						// check if something already exists?
						prayers[int(*inline.Meta.MaxResults)*(page-1)+i] = prayer
					}
				} else {
					return fmt.Errorf("unexpected type %T", inline)
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
