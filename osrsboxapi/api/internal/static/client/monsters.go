package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/atye/gosrsbox/osrsboxapi"
)

func (c *client) GetMonstersByName(ctx context.Context, names ...string) ([]osrsboxapi.Monster, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}

	var monsters []osrsboxapi.Monster
	for _, name := range names {
		result, err := c.GetMonstersByQuery(ctx, fmt.Sprintf(`#(wiki_name=="%s")#|#(duplicate==false)#`, name))
		if err != nil {
			return nil, err
		}
		monsters = append(monsters, result...)
	}
	return monsters, nil
}

func (c *client) GetMonstersByQuery(ctx context.Context, query string) ([]osrsboxapi.Monster, error) {
	var monsters []osrsboxapi.Monster
	err := gjsonQuery(ctx, c.monsters, query, &monsters)
	return monsters, err
}

func (c *client) GetMonstersThatDrop(ctx context.Context, items ...string) ([]osrsboxapi.Monster, error) {
	if len(items) == 0 {
		return nil, errors.New("no items provided")
	}

	var monsters []osrsboxapi.Monster
	for _, item := range items {
		result, err := c.GetMonstersByQuery(ctx, fmt.Sprintf(`#(drops.#(name=="%s"))#|#(duplicate==false)#`, item))
		if err != nil {
			return nil, err
		}
		monsters = append(monsters, result...)
	}
	return monsters, nil
}

func (c *client) UpdateMonsters() error {
	monsters, err := c.source.Monsters()
	if err != nil {
		return err
	}
	c.monsters = monsters
	return nil
}
