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

func (c *APIClient) GetMonstersByName(ctx context.Context, names ...string) ([]db.Monster, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	formattedNames := make([]string, len(names))

	for i, name := range names {
		formattedNames[i] = fmt.Sprintf(`"%s"`, name)
	}

	query := fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(formattedNames, ", "))
	return c.GetMonstersByQuery(ctx, query)
}

func (c *APIClient) GetMonstersThatDrop(ctx context.Context, names ...string) ([]db.Monster, error) {
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

func (c *APIClient) GetMonstersByQuery(ctx context.Context, query string) ([]db.Monster, error) {
	apiURL := fmt.Sprintf("%s/%s?where=%s", c.address, monstersEndpoint, url.QueryEscape(query))

	var monstersResp MonstersResponse
	err := c.doRequest(ctx, apiURL, &monstersResp)
	if monstersResp.Error != nil {
		return nil, monstersResp.Error
	}
	if err != nil {
		return nil, err
	}

	monsters := make([]db.Monster, monstersResp.Meta.Total)
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
				var temp MonstersResponse
				err := c.doRequest(ctx, fmt.Sprintf("%s%s", apiURL, url.QueryEscape(fmt.Sprintf("&page=%d", page))), &temp)
				if temp.Error != nil {
					return temp.Error
				}
				if err != nil {
					return err
				}
				for i, monster := range temp.Monsters {
					c.mu.Lock()
					monsters[monstersResp.Meta.MaxResults*(page-1)+i] = monster
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

	return monsters, nil
}
