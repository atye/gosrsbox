package inmemory

import (
	"context"
	"errors"
	"fmt"

	"github.com/atye/gosrsbox/osrsbox/db"
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

func (c *InMemoryClient) GetItemsByQuery(ctx context.Context, query string) ([]db.Item, error) {
	gjResult, err := c.getByQuery(ctx, "items", query)
	if err != nil {
		return nil, err
	}
	if items, ok := gjResult.([]db.Item); ok {
		return items, nil
	}
	return nil, fmt.Errorf("query result %T is not a valid slice of items", gjResult)
}

func (c *InMemoryClient) UpdateItems() error {
	items, err := c.Updater.Items()
	if err != nil {
		return err
	}
	c.Items = items
	return nil
}
