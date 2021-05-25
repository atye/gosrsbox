package client

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/atye/gosrsbox/models"
	"golang.org/x/sync/errgroup"
)

func (c *apiClient) GetMonstersByID(ctx context.Context, ids ...string) ([]models.Monster, error) {
	if len(ids) == 0 {
		return nil, errNoIDs
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *apiClient) GetMonstersByName(ctx context.Context, names ...string) ([]models.Monster, error) {
	if len(names) == 0 {
		return nil, errNoNames
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *apiClient) GetMonstersThatDrop(ctx context.Context, names ...string) ([]models.Monster, error) {
	if len(names) == 0 {
		return nil, errNoNames
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "drops": { "$elemMatch": { "name": { "$in": [%s] } } }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *apiClient) GetMonstersByQuery(ctx context.Context, query string) ([]models.Monster, error) {
	inline, err := c.doMonstersRequest(ctx, c.openAPIClient.MonsterApi.Getmonsters(ctx).Where(query))
	if err != nil {
		return nil, err
	}

	pages := int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
	monsters := make([]models.Monster, *inline.Meta.Total)

	var tmpMonsters []models.Monster
	err = convert(inline.GetItems(), &tmpMonsters)
	if err != nil {
		return nil, err
	}

	_ = copy(monsters, tmpMonsters)

	if pages > 1 {
		var eg errgroup.Group
		for page := 2; page <= pages; page++ {
			page := page
			eg.Go(func() error {
				inline, err := c.doMonstersRequest(ctx, c.openAPIClient.MonsterApi.Getmonsters(ctx).Where(query).Page(int32(page)))
				if err != nil {
					return err
				}

				var tmpMonsters []models.Monster
				err = convert(inline.GetItems(), &tmpMonsters)
				if err != nil {
					return err
				}

				for i, monster := range tmpMonsters {
					monsters[int(*inline.Meta.MaxResults)*(page-1)+i] = monster
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
