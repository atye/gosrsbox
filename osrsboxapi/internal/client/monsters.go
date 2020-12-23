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

func (c *client) GetMonstersByID(ctx context.Context, ids ...string) ([]openapi.Monster, error) {
	if len(ids) == 0 {
		return nil, errors.New("no ids provided")
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *client) GetMonstersByName(ctx context.Context, names ...string) ([]openapi.Monster, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *client) GetMonstersThatDrop(ctx context.Context, names ...string) ([]openapi.Monster, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "drops": { "$elemMatch": { "name": { "$in": [%s] } } }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *client) GetMonstersByQuery(ctx context.Context, query string) ([]openapi.Monster, error) {
	resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.MonsterApi.Getmonsters(ctx).Where(query))
	if err != nil {
		return nil, err
	}
	switch inline := resp.(type) {
	case openapi.InlineResponse2003:
		pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
		monsters := make([]openapi.Monster, *inline.Meta.Total)
		for i, monster := range inline.GetItems() {
			monsters[i] = monster
		}
		if pages > 1 {
			var eg errgroup.Group
			for page := 2; page <= pages; page++ {
				page := page
				eg.Go(func() error {
					resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.MonsterApi.Getmonsters(ctx).Where(query).Page(int32(page)))
					if err != nil {
						return err
					}
					if inline, ok := resp.(openapi.InlineResponse2003); ok {
						for i, monster := range inline.GetItems() {
							// check if something already exists?
							monsters[int(*inline.Meta.MaxResults)*(page-1)+i] = monster
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
		return monsters, nil
	default:
		return nil, fmt.Errorf("unexpected response type %T", inline)
	}
}
