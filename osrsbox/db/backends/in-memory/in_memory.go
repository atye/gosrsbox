package inmemory

import (
	"context"
	"encoding/json"

	"github.com/tidwall/gjson"
)

type source interface {
	Items() ([]byte, error)
	Monsters() ([]byte, error)
	Prayers() ([]byte, error)
}

type InMemoryClient struct {
	items    []byte
	monsters []byte
	prayers  []byte
	source   source
}

func NewAPI() *InMemoryClient {
	return &InMemoryClient{}
}

func (c *InMemoryClient) RunOptions(options ...Option) error {
	for _, opt := range options {
		err := opt(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func gjsonQuery(ctx context.Context, data []byte, query string, v interface{}) error {
	gjResult := gjson.GetBytes(data, query)
	err := json.Unmarshal([]byte(gjResult.String()), &v)
	if err != nil {
		return err
	}
	return nil
}
