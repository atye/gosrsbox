package client

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/atye/gosrsbox/models"
	"golang.org/x/sync/errgroup"
)

func (c *apiClient) GetPrayersByID(ctx context.Context, ids ...string) ([]models.Prayer, error) {
	if len(ids) == 0 {
		return nil, errors.New("no ids provided")
	}
	return c.GetPrayersByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }}`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *apiClient) GetPrayersByName(ctx context.Context, names ...string) ([]models.Prayer, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}

	query := fmt.Sprintf(`{ "name": { "$in": [%s] } }`, strings.Join(quoteStrings(names...), ", "))
	return c.GetPrayersByQuery(ctx, query)
}

func (c *apiClient) GetPrayersByQuery(ctx context.Context, query string) ([]models.Prayer, error) {
	inline, err := c.doPrayersRequest(ctx, c.openAPIClient.PrayerApi.Getprayers(ctx).Where(query))
	if err != nil {
		return nil, err
	}

	pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
	prayers := make([]models.Prayer, *inline.Meta.Total)

	/*var tmpPrayers []models.Prayer
	err = convert(inline.GetItems(), &tmpPrayers)
	if err != nil {
		return nil, err
	}*/

	_ = copy(prayers, inline.GetItems())

	if pages > 1 {
		var eg errgroup.Group
		for page := 2; page <= pages; page++ {
			page := page
			eg.Go(func() error {
				inline, err := c.doPrayersRequest(ctx, c.openAPIClient.PrayerApi.Getprayers(ctx).Where(query).Page(int32(page)))
				if err != nil {
					return err
				}

				/*var tmpPrayers []models.Prayer
				err = convert(inline.GetItems(), &tmpPrayers)
				if err != nil {
					return err
				}*/

				for i, prayer := range inline.GetItems() {
					prayers[int(*inline.Meta.MaxResults)*(page-1)+i] = prayer
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
