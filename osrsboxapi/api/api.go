package api

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/atye/gosrsbox/osrsboxapi/api/internal/client"
	"github.com/atye/gosrsbox/osrsboxapi/api/internal/client/openapi"

	"github.com/atye/gosrsbox/osrsboxapi/sets"
	"github.com/atye/gosrsbox/osrsboxapi/slots"
)

type API interface {
	// GetItemsByName returns a slice of Items from the given wiki names
	GetItemsByName(ctx context.Context, names ...string) ([]openapi.Item, error)

	// GetItemsByQuery returns a slice of Items from the given MongoDB or Python query
	// ex:
	// api.GetItemsByQuery(context.Background(), "equipment.prayer<0")
	// api.GetItemsByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	GetItemsByQuery(ctx context.Context, query string) ([]openapi.Item, error)

	// GetItemSet returns a slice of Items in the given set
	GetItemSet(ctx context.Context, set sets.SetName) ([]openapi.Item, error)

	// GetItemsBySlot returns a slice of Items in the given slot
	GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]openapi.Item, error)

	// GetMonstersByName returns a slice of Monsters from the given wiki names
	GetMonstersByName(ctx context.Context, names ...string) ([]openapi.Monster, error)

	// GetMonstersByQuery returns a slice of Monsters from the given MongoDB or Python
	// ex:
	// api.GetMonstersByQuery(context.Background(), "equipment.prayer<0")
	// api.GetMonstersByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	GetMonstersByQuery(ctx context.Context, query string) ([]openapi.Monster, error)

	// GetMonstersThatDrop returns a slice of Monsters that drop the given items
	GetMonstersThatDrop(ctx context.Context, items ...string) ([]openapi.Monster, error)

	// GetPrayersByName returns a slice of Prayers from the given names
	GetPrayersByName(ctx context.Context, names ...string) ([]openapi.Prayer, error)

	// GetPrayersByQuery returns a slice of Prayers form the given MongoDB or Python query
	// ex:
	// api.GetPrayersByQuery(context.Background(), "equipment.prayer<0")
	// api.GetPrayersByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	GetPrayersByQuery(ctx context.Context, query string) ([]openapi.Prayer, error)

	// GetJSONFile retrieves the specified JSON files from the Static JSON API and unmarshals into an interfaces of your choosing
	// ex:
	// var twoHandedITems map[string]osrsboxapi.Item
	// err := api.GetJSONFiles(context.Background(), []string{"items-json-slot/items-2h.json"}, &twoHandedITems)
	GetJSONFiles(ctx context.Context, files []string, destinations ...interface{}) error
}

type APIConfig struct {
	Logger     *log.Logger
	HttpClient *http.Client
}

var (
	once sync.Once
	api  API
)

func NewAPI(config *APIConfig) API {
	once.Do(func() {
		logger, httpClient := logger(config), httpClient(config)
		conf := &openapi.Configuration{
			Scheme:     "https",
			HTTPClient: httpClient,
			Servers: []openapi.ServerConfiguration{
				{
					URL: "api.osrsbox.com",
				},
			},
		}
		api = withLogger(client.NewAPI(conf), logger)
	})
	return api
}

func logger(c *APIConfig) *log.Logger {
	var logger *log.Logger
	if c == nil {
		logger = log.New(os.Stdout, "osrsboxapi", log.LstdFlags)
	} else if c.Logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	} else {
		logger = c.Logger
	}
	return logger
}

func httpClient(c *APIConfig) *http.Client {
	if c != nil && c.HttpClient != nil {
		return c.HttpClient
	}
	return http.DefaultClient
}
