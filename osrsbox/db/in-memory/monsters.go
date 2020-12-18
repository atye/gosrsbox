package inmemory

import (
	"context"
	"errors"
	"fmt"

	"github.com/atye/gosrsbox/osrsbox/db"
)

func (c *InMemoryClient) GetMonstersByName(ctx context.Context, names ...string) ([]db.Monster, error) {
	if len(names) == 0 {
		return nil, errors.New("No names provided")
	}

	var monsters []db.Monster
	for _, name := range names {
		result, err := c.GetMonstersByQuery(ctx, fmt.Sprintf(`#(wiki_name=="%s")#|#(duplicate==false)#`, name))
		if err != nil {
			return nil, err
		}

		monsters = append(monsters, result...)
	}
	return monsters, nil
}

func (c *InMemoryClient) GetMonstersByQuery(ctx context.Context, query string) ([]db.Monster, error) {
	gjResult, err := c.getByQuery(ctx, "monsters", query)
	if err != nil {
		return nil, err
	}
	if monsters, ok := gjResult.([]db.Monster); ok {
		return monsters, nil
	}
	return nil, fmt.Errorf("query result %T is not a valid slice of items", gjResult)
}

func (c *InMemoryClient) GetMonstersThatDrop(ctx context.Context, items ...string) ([]db.Monster, error) {
	var monsters []db.Monster
	for _, item := range items {
		gjResult, err := c.getByQuery(ctx, "monsters", fmt.Sprintf(`#(drops.#(name=="%s"))#|#(duplicate==false)#`, item))
		if err != nil {
			return nil, err
		}

		if result, ok := gjResult.([]db.Monster); ok {
			monsters = append(monsters, result...)
		}
	}
	return monsters, nil
}

func (c *InMemoryClient) UpdateMonsters() error {
	monsters, err := c.Updater.Monsters()
	if err != nil {
		return err
	}
	c.Monsters = monsters
	return nil
}
