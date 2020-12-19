package inmemory

import (
	"context"
	"errors"
	"fmt"

	"github.com/atye/gosrsbox/osrsbox/db"
)

func (c *InMemoryClient) GetMonstersByName(ctx context.Context, names ...string) ([]db.Monster, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
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
	var monsters []db.Monster
	err := gjsonQuery(ctx, c.monsters, query, &monsters)
	return monsters, err
}

func (c *InMemoryClient) GetMonstersThatDrop(ctx context.Context, items ...string) ([]db.Monster, error) {
	if len(items) == 0 {
		return nil, errors.New("no items provided")
	}

	var monsters []db.Monster
	for _, item := range items {
		result, err := c.GetMonstersByQuery(ctx, fmt.Sprintf(`#(drops.#(name=="%s"))#|#(duplicate==false)#`, item))
		if err != nil {
			return nil, err
		}
		monsters = append(monsters, result...)
		/*gjResult, err := c.getByQuery(ctx, "monsters", fmt.Sprintf(`#(drops.#(name=="%s"))#|#(duplicate==false)#`, item))
		if err != nil {
			return nil, err
		}

		if result, ok := gjResult.([]db.Monster); ok {
			monsters = append(monsters, result...)
		}*/
	}
	return monsters, nil
}

func (c *InMemoryClient) UpdateMonsters() error {
	monsters, err := c.source.Monsters()
	if err != nil {
		return err
	}
	c.monsters = monsters
	return nil
}
