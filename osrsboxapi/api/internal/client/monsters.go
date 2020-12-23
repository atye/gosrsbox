package client

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/atye/gosrsbox/osrsboxapi/api/internal/client/openapi"
	"golang.org/x/sync/errgroup"
)

func (c *client) GetMonstersByName(ctx context.Context, names ...string) ([]openapi.Monster, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	query := fmt.Sprintf(`{ "wiki_name": { "$in": [%s] }, "duplicate": false }`, strings.Join(quoteStrings(names...), ", "))
	return c.GetMonstersByQuery(ctx, query)
}

func (c *client) GetMonstersThatDrop(ctx context.Context, names ...string) ([]openapi.Monster, error) {
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

func (c *client) GetMonstersByQuery(ctx context.Context, query string) ([]openapi.Monster, error) {
	resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.Getmonsters(ctx).Where(query))
	if err != nil {
		return nil, err
	}
	var pages int
	var inline openapi.InlineResponse2003
	if inline, ok := resp.(openapi.InlineResponse2003); ok {
		pages = int(math.Ceil(float64(*inline.Meta.Total) / float64(*inline.Meta.MaxResults)))
	} else {
		return nil, fmt.Errorf("%T", inline)
	}
	monsters := make([]openapi.Monster, *inline.Meta.Total)
	for i, monster := range inline.GetItems() {
		monsters[i] = monster
	}
	if pages > 1 {
		var eg errgroup.Group
		for page := 2; page <= pages; page++ {
			page := page
			eg.Go(func() error {
				resp, err := c.doOpenAPIRequest(ctx, c.openAPIClient.Getmonsters(ctx).Where(query).Page(int32(page)))
				if err != nil {
					return err
				}
				if inline, ok := resp.(openapi.InlineResponse2003); ok {
					for i, monster := range *inline.Items {
						// check if something already exists?
						monsters[int(*inline.Meta.MaxResults)*(page-1)+i] = monster
					}
				} else {
					return fmt.Errorf("%T", inline)
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
