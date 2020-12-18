package inmemory

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/atye/gosrsbox/osrsbox/db"
	"github.com/tidwall/gjson"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type Updater interface {
	Items() ([]byte, error)
	Monsters() ([]byte, error)
	Prayers() ([]byte, error)
}

type InMemoryClient struct {
	Items    []byte
	Monsters []byte
	Prayers  []byte
	Updater  Updater
	Logger   Logger
}

func NewInMemoryClient(options ...Option) (*InMemoryClient, error) {
	db := &InMemoryClient{}
	for _, opt := range options {
		err := opt(db)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func (c *InMemoryClient) getByQuery(ctx context.Context, data, query string) (interface{}, error) {
	switch data {
	case "items":
		gjResult := gjson.GetBytes(c.Items, query)
		var items []db.Item
		err := json.Unmarshal([]byte(gjResult.String()), &items)
		if err != nil {
			return nil, err
		}
		return items, err
	case "monsters":
		gjResult := gjson.GetBytes(c.Monsters, query)
		var monsters []db.Monster
		err := json.Unmarshal([]byte(gjResult.String()), &monsters)
		if err != nil {
			return nil, err
		}
		return monsters, err
	case "prayers":
		gjResult := gjson.GetBytes(c.Prayers, query)
		var prayers []db.Prayer
		err := json.Unmarshal([]byte(gjResult.String()), &prayers)
		if err != nil {
			return nil, err
		}
		return prayers, err
	default:
		return nil, fmt.Errorf("%s not supported", data)
	}
}
