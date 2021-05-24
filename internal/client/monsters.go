package client

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/atye/gosrsbox/openapi/api"
	"golang.org/x/sync/errgroup"
)

func (c *client) GetMonstersByID(ctx context.Context, ids ...string) ([]api.Monster, error) {
	if len(ids) == 0 {
		return nil, errNoIDs
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *client) GetMonstersByName(ctx context.Context, names ...string) ([]api.Monster, error) {
	if len(names) == 0 {
		return nil, errNoNames
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *client) GetMonstersThatDrop(ctx context.Context, names ...string) ([]api.Monster, error) {
	if len(names) == 0 {
		return nil, errNoNames
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "drops": { "$elemMatch": { "name": { "$in": [%s] } } }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *client) GetMonstersByQuery(ctx context.Context, query string) ([]api.Monster, error) {
	resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.MonsterApi.Getmonsters(ctx).Where(query))
	if err != nil {
		return nil, err
	}

	switch inline := resp.(type) {
	case api.InlineResponse2003:
		pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
		monsters := make([]api.Monster, *inline.Meta.Total)

		_ = copy(monsters, inline.GetItems())

		if pages > 1 {
			var eg errgroup.Group
			for page := 2; page <= pages; page++ {
				page := page
				eg.Go(func() error {
					resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.MonsterApi.Getmonsters(ctx).Where(query).Page(int32(page)))
					if err != nil {
						return err
					}

					switch inline := resp.(type) {
					case api.InlineResponse2003:
						for i, monster := range inline.GetItems() {
							monsters[int(*inline.Meta.MaxResults)*(page-1)+i] = monster
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
		return monsters, nil
	default:
		return nil, fmt.Errorf("unexpected response type %T", inline)
	}
}
