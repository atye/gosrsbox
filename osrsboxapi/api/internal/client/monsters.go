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

func (c *client) GetMonstersByName(ctx context.Context, names ...string) ([]osrsboxapi.Monster, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	query := fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", "))
	return c.GetMonstersByQuery(ctx, query)
}

func (c *client) GetMonstersThatDrop(ctx context.Context, names ...string) ([]osrsboxapi.Monster, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	formattedNames := make([]string, len(names))

	for i, name := range names {
		formattedNames[i] = fmt.Sprintf(`"%s"`, name)
	}

	query := fmt.Sprintf(`{ "drops": { "$elemMatch": { "name": { "$in": [%s] } } }, "duplicate": false }`, strings.Join(formattedNames, ", "))
	return c.GetMonstersByQuery(ctx, query)
}

func (c *client) GetMonstersByQuery(ctx context.Context, query string) ([]osrsboxapi.Monster, error) {
	apiURL := fmt.Sprintf("%s/%s?where=%s", c.apiAddress, monstersEndpoint, url.QueryEscape(query))

	var monstersResp monstersResponse
	_, err := c.doAPIRequest(ctx, apiURL, &monstersResp)
	if monstersResp.Error != nil {
		return nil, monstersResp.Error
	}
	if err != nil {
		return nil, err
	}

	monsters := make([]osrsboxapi.Monster, monstersResp.Meta.Total)
	for i, monster := range monstersResp.Monsters {
		monsters[i] = monster
	}

	var pages int
	if monstersResp.Meta.MaxResults != 0 {
		pages = int(math.Ceil(float64(monstersResp.Meta.Total) / float64(monstersResp.Meta.MaxResults)))
	}

	if pages > 1 {
		var eg errgroup.Group
		for page := 2; page <= pages; page++ {
			page := page
			eg.Go(func() error {
				var temp monstersResponse
				_, err := c.doAPIRequest(ctx, fmt.Sprintf("%s%s", apiURL, fmt.Sprintf("&page=%d", page)), &temp)
				if temp.Error != nil {
					return temp.Error
				}
				if err != nil {
					return err
				}
				for i, monster := range temp.Monsters {
					monsters[temp.Meta.MaxResults*(page-1)+i] = monster
				}
				return nil
			})
		}
		err := eg.Wait()
		if err != nil {
			return nil, err
		}
	}

	return monsters, nil
}
