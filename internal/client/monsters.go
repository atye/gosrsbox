package client

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/atye/gosrsbox/internal/common"
	"github.com/atye/gosrsbox/models"
	"golang.org/x/sync/errgroup"
)

func (c *APIClient) GetMonstersByID(ctx context.Context, ids ...string) ([]models.Monster, error) {
	ctx, span := c.createSpan(ctx, "get_monsters_by_id")
	defer span.End()

	if len(ids) == 0 {
		return nil, errNoIDs
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
}

func (c *APIClient) GetMonstersByName(ctx context.Context, names ...string) ([]models.Monster, error) {
	ctx, span := c.createSpan(ctx, "get_monsters_by_name")
	defer span.End()

	if len(names) == 0 {
		return nil, errNoNames
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *APIClient) GetMonstersThatDrop(ctx context.Context, names ...string) ([]models.Monster, error) {
	ctx, span := c.createSpan(ctx, "get_monsters_that_drop")
	defer span.End()

	if len(names) == 0 {
		return nil, errNoNames
	}
	return c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "drops": { "$elemMatch": { "name": { "$in": [%s] } } }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
}

func (c *APIClient) GetMonstersByQuery(ctx context.Context, query string) ([]models.Monster, error) {
	ctx, span := c.createSpan(ctx, "get_monsters_by_query")
	defer span.End()

	inline, err := c.doMonstersRequest(ctx, common.Params{Where: query})
	if err != nil {
		return nil, err
	}

	total := inline.GetTotal()
	maxResults := inline.GetMaxResults()

	pages := int(math.Ceil(float64(total) / float64(maxResults)))
	monsters := make([]models.Monster, total)

	_ = copy(monsters, inline.GetMonsters())

	if pages > 1 {
		var eg errgroup.Group
		for page := 2; page <= pages; page++ {
			page := page
			eg.Go(func() error {
				inline, err := c.doMonstersRequest(ctx, common.Params{Where: query, Page: page})
				if err != nil {
					return err
				}

				for i, monster := range inline.GetMonsters() {
					monsters[maxResults*(page-1)+i] = monster
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
