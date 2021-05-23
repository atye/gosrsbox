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

func (c *client) GetMonstersByID(ctx context.Context, ids ...string) ([]osrsbox.Monster, error) {
	if len(ids) == 0 {
		return nil, errors.New("no ids provided")
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *client) GetMonstersByName(ctx context.Context, names ...string) ([]osrsbox.Monster, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *client) GetMonstersThatDrop(ctx context.Context, names ...string) ([]osrsbox.Monster, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "drops": { "$elemMatch": { "name": { "$in": [%s] } } }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *client) GetMonstersByQuery(ctx context.Context, query string) ([]osrsbox.Monster, error) {
	resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.MonsterApi.Getmonsters(ctx).Where(query))
	if err != nil {
		return nil, err
	}

	switch inline := resp.(type) {
	case openapi.InlineResponse2003:
		pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
		monsters := make([]osrsbox.Monster, *inline.Meta.Total)

		respMonsters, err := openMonstersToMonsters(inline.GetItems())
		if err != nil {
			return nil, err
		}

		_ = copy(monsters, respMonsters)

		if pages > 1 {
			var eg errgroup.Group
			for page := 2; page <= pages; page++ {
				page := page
				eg.Go(func() error {
					resp, err := c.retryOpenAPIRequest(ctx, c.openAPIClient.MonsterApi.Getmonsters(ctx).Where(query).Page(int32(page)), errTooManyOpenFiles)
					if err != nil {
						return err
					}

					switch inline := resp.(type) {
					case openapi.InlineResponse2003:
						respMonsters, err := openMonstersToMonsters(inline.GetItems())
						if err != nil {
							return err
						}

						for i, monster := range respMonsters {
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

func openMonstersToMonsters(openMonsters []openapi.Monster) ([]osrsbox.Monster, error) {
	b, err := json.Marshal(openMonsters)
	if err != nil {
		return nil, err
	}

	var monsters []osrsbox.Monster
	err = json.Unmarshal(b, &monsters)
	if err != nil {
		return nil, err
	}

	return monsters, nil
}
