package client

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/url"
	"strings"

	"github.com/atye/gosrsbox/osrsboxdb"
	"golang.org/x/sync/errgroup"
)

func (c *client) GetPrayersByName(ctx context.Context, names ...string) ([]osrsboxdb.Prayer, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	query := fmt.Sprintf(`{ "name": { "$in": [%s] } }`, strings.Join(quoteStrings(names...), ", "))
	return c.GetPrayersByQuery(ctx, query)
}

func (c *client) GetPrayersByQuery(ctx context.Context, query string) ([]osrsboxdb.Prayer, error) {
	apiURL := fmt.Sprintf("%s/%s?where=%s", c.address, prayersEndpoint, url.QueryEscape(query))

	var prayersResp prayersResponse
	err := c.doRequest(ctx, apiURL, &prayersResp)
	if prayersResp.Error != nil {
		return nil, prayersResp.Error
	}
	if err != nil {
		return nil, err
	}

	prayers := make([]osrsboxdb.Prayer, prayersResp.Meta.Total)
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
				err := c.doRequest(ctx, fmt.Sprintf("%s%s", apiURL, fmt.Sprintf("&page=%d", page)), &temp)
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
