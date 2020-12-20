package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/atye/gosrsbox/osrsboxdb"
	"github.com/atye/gosrsbox/osrsboxdb/sets"
)

func (c *client) GetItemsByName(ctx context.Context, names ...string) ([]osrsboxdb.Item, error) {
	var items []osrsboxdb.Item

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

func (c *client) GetItemSet(ctx context.Context, setName sets.SetName) ([]osrsboxdb.Item, error) {
	if setName == nil || len(setName) == 0 {
		return nil, errors.New("no set provided")
	}
	return c.GetItemsByName(ctx, setName...)
}

func (c *client) GetItemsByQuery(ctx context.Context, query string) ([]osrsboxdb.Item, error) {
	var items []osrsboxdb.Item
	err := gjsonQuery(ctx, c.items, query, &items)
	return items, err
}

func (c *client) UpdateItems() error {
	items, err := c.source.Items()
	if err != nil {
		return err
	}
	c.items = items
	return nil
}
