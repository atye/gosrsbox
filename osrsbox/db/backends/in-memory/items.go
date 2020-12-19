package inmemory

import (
	"context"
	"errors"
	"fmt"

	"github.com/atye/gosrsbox/osrsbox/db"
	"github.com/atye/gosrsbox/osrsbox/db/sets"
)

func (c *InMemoryClient) GetItemsByName(ctx context.Context, names ...string) ([]db.Item, error) {
	var items []db.Item

	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	for _, name := range names {
		result, err := c.GetItemsByQuery(ctx, fmt.Sprintf(`#(wiki_name=="%s")#|#(duplicate==false)#`, name))
		if err != nil {
			return nil, err
		}
		items = append(items, result...)
	}
	return items, nil
}

func (c *InMemoryClient) GetItemSet(ctx context.Context, setName sets.SetName) ([]db.Item, error) {
	if setName == nil || len(setName) == 0 {
		return nil, errors.New("no set provided")
	}
	return c.GetItemsByName(ctx, setName...)
}

func (c *InMemoryClient) GetItemsByQuery(ctx context.Context, query string) ([]db.Item, error) {
	var items []db.Item
	err := gjsonQuery(ctx, c.items, query, &items)
	return items, err
}

func (c *InMemoryClient) UpdateItems() error {
	items, err := c.source.Items()
	if err != nil {
		return err
	}
	c.items = items
	return nil
}
