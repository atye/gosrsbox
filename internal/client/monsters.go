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

	var err error
	defer func() {
		if err != nil {
			setSpanErrorStatus(span, err)
		}
	}()

	if len(ids) == 0 {
		err = errNoIDs
		return nil, err
	}
	monsters, err := c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "id": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(ids...), ", ")))
	if err != nil {
		return nil, err
	}
	return monsters, nil
}

func (c *APIClient) GetMonstersByName(ctx context.Context, names ...string) ([]models.Monster, error) {
	ctx, span := c.createSpan(ctx, "get_monsters_by_name")
	defer span.End()

	var err error
	defer func() {
		if err != nil {
			setSpanErrorStatus(span, err)
		}
	}()

	if len(names) == 0 {
		err = errNoNames
		return nil, err
	}
	monsters, err := c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
	if err != nil {
		return nil, err
	}
	return monsters, nil
}

func (c *APIClient) GetMonstersThatDrop(ctx context.Context, names ...string) ([]models.Monster, error) {
	ctx, span := c.createSpan(ctx, "get_monsters_that_drop")
	defer span.End()

	var err error
	defer func() {
		if err != nil {
			setSpanErrorStatus(span, err)
		}
	}()

	if len(names) == 0 {
		err = errNoNames
		return nil, err
	}
	monsters, err := c.GetMonstersByQuery(ctx, fmt.Sprintf(`{ "drops": { "$elemMatch": { "name": { "$in": [%s] } } }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", ")))
	if err != nil {
		return nil, err
	}
	return monsters, nil
}

func (c *APIClient) GetMonstersByQuery(ctx context.Context, query string) ([]models.Monster, error) {
	ctx, span := c.createSpan(ctx, "get_monsters_by_query")
	defer span.End()

	var err error
	defer func() {
		if err != nil {
			setSpanErrorStatus(span, err)
		}
	}()

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
		err = eg.Wait()
		if err != nil {
			return nil, err
		}
	}
	return monsters, nil
}
