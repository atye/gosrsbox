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
	"golang.org/x/sync/errgroup"
)

func (c *client) GetPrayersByID(ctx context.Context, ids ...string) ([]osrsbox.Prayer, error) {
	if len(ids) == 0 {
		return nil, errors.New("no ids provided")
	}
	return c.GetPrayersByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }}`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *client) GetPrayersByName(ctx context.Context, names ...string) ([]osrsbox.Prayer, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}

	query := fmt.Sprintf(`{ "name": { "$in": [%s] } }`, strings.Join(quoteStrings(names...), ", "))
	return c.GetPrayersByQuery(ctx, query)
}

func (c *client) GetPrayersByQuery(ctx context.Context, query string) ([]osrsbox.Prayer, error) {
	resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.PrayerApi.Getprayers(ctx).Where(query))
	if err != nil {
		return nil, err
	}

	switch inline := resp.(type) {
	case openapi.InlineResponse2004:
		pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
		prayers := make([]osrsbox.Prayer, *inline.Meta.Total)

		respPrayers, err := openPrayersToPrayers(inline.GetItems())
		if err != nil {
			return nil, err
		}

		_ = copy(prayers, respPrayers)

		if pages > 1 {
			var eg errgroup.Group
			for page := 2; page <= pages; page++ {
				page := page
				eg.Go(func() error {
					resp, err := c.retryOpenAPIRequest(ctx, c.openAPIClient.PrayerApi.Getprayers(ctx).Where(query).Page(int32(page)), errTooManyOpenFiles)
					if err != nil {
						return err
					}

					switch inline := resp.(type) {
					case openapi.InlineResponse2004:
						respPrayers, err := openPrayersToPrayers(inline.GetItems())
						if err != nil {
							return err
						}

						for i, prayer := range respPrayers {
							prayers[int(*inline.Meta.MaxResults)*(page-1)+i] = prayer
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
		return prayers, nil
	default:
		return nil, fmt.Errorf("unexpected response type %T", inline)
	}
}

func openPrayersToPrayers(openPrayers []openapi.Prayer) ([]osrsbox.Prayer, error) {
	b, err := json.Marshal(openPrayers)
	if err != nil {
		return nil, err
	}

	var prayers []osrsbox.Prayer
	err = json.Unmarshal(b, &prayers)
	if err != nil {
		return nil, err
	}

	return prayers, nil
}
