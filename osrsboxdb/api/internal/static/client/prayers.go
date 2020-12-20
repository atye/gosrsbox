package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/atye/gosrsbox/osrsboxdb"
)

func (c *client) GetPrayersByName(ctx context.Context, names ...string) ([]osrsboxdb.Prayer, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	var prayers []osrsboxdb.Prayer
	for _, name := range names {
		result, err := c.GetPrayersByQuery(ctx, fmt.Sprintf(`#(name=="%s")#`, name))
		if err != nil {
			return nil, err
		}
		prayers = append(prayers, result...)
	}
	return prayers, nil
}

func (c *client) GetPrayersByQuery(ctx context.Context, query string) ([]osrsboxdb.Prayer, error) {
	var prayers []osrsboxdb.Prayer
	err := gjsonQuery(ctx, c.prayers, query, &prayers)
	return prayers, err
}

func (c *client) UpdatePrayers() error {
	prayers, err := c.source.Prayers()
	if err != nil {
		return err
	}
	c.prayers = prayers
	return nil
}
