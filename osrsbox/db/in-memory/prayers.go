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
	gjResult, err := c.getByQuery(ctx, "prayers", query)
	if err != nil {
		return nil, err
	}
	if prayers, ok := gjResult.([]db.Prayer); ok {
		return prayers, nil
	}
	return nil, fmt.Errorf("query result %T is not a valid slice of prayers", gjResult)
}

func (c *InMemoryClient) UpdatePrayers() error {
	prayers, err := c.Updater.Prayers()
	if err != nil {
		return err
	}
	c.Prayers = prayers
	return nil
}
