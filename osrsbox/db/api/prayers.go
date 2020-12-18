package api

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/url"
	"strings"

	"github.com/atye/gosrsbox/osrsbox/db"
	"golang.org/x/sync/errgroup"
)

func (c *APIClient) GetPrayersByName(ctx context.Context, names ...string) ([]db.Prayer, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	formattedNames := make([]string, len(names))

	for i, name := range names {
		formattedNames[i] = fmt.Sprintf(`"%s"`, name)
	}

	query := fmt.Sprintf(`{ "name": { "$in": [%s] } }`, strings.Join(formattedNames, ", "))
	return c.GetPrayersByQuery(ctx, query)
}

func (c *APIClient) GetPrayersByQuery(ctx context.Context, query string) ([]db.Prayer, error) {
	apiURL := fmt.Sprintf("%s/%s?where=%s", c.address, prayersEndpoint, url.QueryEscape(query))

	var prayersResp PrayersResponse
	err := c.doRequest(ctx, apiURL, &prayersResp)
	if prayersResp.Error != nil {
		return nil, prayersResp.Error
	}
	if err != nil {
		return nil, err
	}

	prayers := make([]db.Prayer, prayersResp.Meta.Total)
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
				var temp PrayersResponse
				err := c.doRequest(ctx, fmt.Sprintf("%s%s", apiURL, url.QueryEscape(fmt.Sprintf("&page=%d", page))), &temp)
				if temp.Error != nil {
					return temp.Error
				}
				if err != nil {
					return err
				}
				for i, prayer := range temp.Prayers {
					c.mu.Lock()
					prayers[prayersResp.Meta.MaxResults*(page-1)+i] = prayer
					c.mu.Unlock()
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
