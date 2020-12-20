package client

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/url"
	"strings"

	"github.com/atye/gosrsbox/osrsboxapi"
	"golang.org/x/sync/errgroup"
)

func (c *client) GetPrayersByName(ctx context.Context, names ...string) ([]osrsboxapi.Prayer, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	query := fmt.Sprintf(`{ "name": { "$in": [%s] } }`, strings.Join(quoteStrings(names...), ", "))
	return c.GetPrayersByQuery(ctx, query)
}

func (c *client) GetPrayersByQuery(ctx context.Context, query string) ([]osrsboxapi.Prayer, error) {
	apiURL := fmt.Sprintf("%s/%s?where=%s", c.apiAddress, prayersEndpoint, url.QueryEscape(query))

	var prayersResp prayersResponse
	_, err := c.doAPIRequest(ctx, apiURL, &prayersResp)
	if prayersResp.Error != nil {
		return nil, prayersResp.Error
	}
	if err != nil {
		return nil, err
	}

	prayers := make([]osrsboxapi.Prayer, prayersResp.Meta.Total)
	for i, prayer := range prayersResp.Prayers {
		prayers[i] = prayer
	}

	var pages int
	if prayersResp.Meta.MaxResults != 0 {
		pages = int(math.Ceil(float64(prayersResp.Meta.Total) / float64(prayersResp.Meta.MaxResults)))
	}

	if pages > 1 {
		var eg errgroup.Group
		for page := 2; page <= pages; page++ {
			page := page
			eg.Go(func() error {
				var temp prayersResponse
				_, err := c.doAPIRequest(ctx, fmt.Sprintf("%s%s", apiURL, fmt.Sprintf("&page=%d", page)), &temp)
				if temp.Error != nil {
					return temp.Error
				}
				if err != nil {
					return err
				}
				for i, prayer := range temp.Prayers {
					prayers[temp.Meta.MaxResults*(page-1)+i] = prayer
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
