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

func (c *APIClient) GetItemsByName(ctx context.Context, names ...string) ([]db.Item, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	formattedNames := make([]string, len(names))

	for i, name := range names {
		formattedNames[i] = fmt.Sprintf(`"%s"`, name)
	}

	query := fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(formattedNames, ", "))
	return c.GetItemsByQuery(ctx, query)
}

func (c *APIClient) GetItemsByQuery(ctx context.Context, query string) ([]db.Item, error) {
	apiURL := fmt.Sprintf("%s/%s?where=%s", c.address, itemsEndpoint, url.QueryEscape(query))

	var itemsResp ItemsResponse
	err := c.doRequest(ctx, apiURL, &itemsResp)
	if itemsResp.Error != nil {
		return nil, itemsResp.Error
	}
	if err != nil {
		return nil, err
	}

	items := make([]db.Item, itemsResp.Meta.Total)
	for i, item := range itemsResp.Items {
		items[i] = item
	}

	var pages int
	if itemsResp.Meta.MaxResults != 0 {
		pages = int(math.Ceil(float64(itemsResp.Meta.Total) / float64(itemsResp.Meta.MaxResults)))
	}

	if pages > 1 {
		var eg errgroup.Group
		for page := 2; page <= pages; page++ {
			page := page
			eg.Go(func() error {
				var temp ItemsResponse
				err := c.doRequest(ctx, fmt.Sprintf("%s%s", apiURL, url.QueryEscape(fmt.Sprintf("&page=%d", page))), &temp)
				if temp.Error != nil {
					return temp.Error
				}
				if err != nil {
					return err
				}
				for i, item := range temp.Items {
					//c.mu.Lock()
					items[itemsResp.Meta.MaxResults*(page-1)+i] = item
					//c.mu.Unlock()
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
