package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/tidwall/gjson"
	"golang.org/x/sync/semaphore"
)

type source interface {
	Items() ([]byte, error)
	Monsters() ([]byte, error)
	Prayers() ([]byte, error)
}

type client struct {
	items    []byte
	monsters []byte
	prayers  []byte
	source   source
	client   *http.Client
	sem      *semaphore.Weighted
	address  string
}

func NewAPI(httpClient *http.Client) *client {
	return &client{
		client: httpClient,
		//address: api,
		sem: semaphore.NewWeighted(int64(10)),
	}
}

func (c *client) RunOptions(options ...Option) error {
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
	if gjResult.Exists() {
		err := json.Unmarshal([]byte(gjResult.Raw), &v)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

/*func gjsonQueryMany(ctx context.Context, data []byte, entity string, queries ...string) (interface{}, error) {
	results := gjson.GetManyBytes(data, queries...)
	switch entity {
	case "monsters":
		var monsters []osrsboxdb.Monster
		for _, result := range results {
			var temp []osrsboxdb.Monster
			err := json.Unmarshal([]byte(result.String()), &temp)
			if err != nil {
				return nil, err
			}
			monsters = append(monsters, temp...)
		}
		return monsters, nil
	case "items":
		var items []osrsboxdb.Item
		for _, result := range results {
			var temp []osrsboxdb.Item
			err := json.Unmarshal([]byte(result.String()), &temp)
			if err != nil {
				return nil, err
			}
			items = append(items, temp...)
		}
		return items, nil
	case "prayers":
		var prayers []osrsboxdb.Prayer
		for _, result := range results {
			var temp []osrsboxdb.Prayer
			err := json.Unmarshal([]byte(result.String()), &temp)
			if err != nil {
				return nil, err
			}
			prayers = append(prayers, temp...)
		}
		return prayers, nil
	default:
		return nil, fmt.Errorf("%s not supported", entity)
	}
}*/
