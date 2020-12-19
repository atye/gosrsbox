package inmemory

import (
	"context"
	"errors"
	"fmt"

	"github.com/atye/gosrsbox/osrsbox/db"
)

func (c *InMemoryClient) GetPrayersByName(ctx context.Context, names ...string) ([]db.Prayer, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	var prayers []db.Prayer
	for _, name := range names {
		result, err := c.GetPrayersByQuery(ctx, fmt.Sprintf(`#(name=="%s")#`, name))
		if err != nil {
			return nil, err
		}
		prayers = append(prayers, result...)
	}
	return prayers, nil
}

func (c *InMemoryClient) GetPrayersByQuery(ctx context.Context, query string) ([]db.Prayer, error) {
	var prayers []db.Prayer
	err := gjsonQuery(ctx, c.prayers, query, &prayers)
	return prayers, err
}

func (c *InMemoryClient) UpdatePrayers() error {
	prayers, err := c.source.Prayers()
	if err != nil {
		return err
	}
	c.prayers = prayers
	return nil
}
