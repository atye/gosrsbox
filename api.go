package gosrsbox

import (
	"context"
	"net/http"
	"sync"

	"github.com/atye/gosrsbox/internal/api"
	openapi "github.com/atye/gosrsbox/internal/openapi/api"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
)

type API interface {
	// GetItemsByName returns a slice of Items with the given IDs
	GetItemsByID(ctx context.Context, ids ...string) ([]api.Item, error)

	// GetItemsByName returns a slice of Items from the given wiki names
	GetItemsByName(ctx context.Context, names ...string) ([]api.Item, error)

	// GetItemsByQuery returns a slice of Items from the given MongoDB or Python query
	// ex:
	// api.GetItemsByQuery(context.Background(), "equipment.prayer<0")
	// api.GetItemsByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	GetItemsByQuery(ctx context.Context, query string) ([]api.Item, error)

	// GetItemSet returns a slice of Items in the given set
	GetItemSet(ctx context.Context, set sets.SetName) ([]api.Item, error)

	// GetItemsBySlot returns a slice of Items in the given slot
	GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]api.Item, error)

	// GetMonstersByID returns a slice of Monsters with the given IDs
	GetMonstersByID(ctx context.Context, ids ...string) ([]openapi.Monster, error)

	// GetMonstersByName returns a slice of Monsters from the given wiki names
	GetMonstersByName(ctx context.Context, names ...string) ([]openapi.Monster, error)

	// GetMonstersByQuery returns a slice of Monsters from the given MongoDB or Python
	// ex:
	// api.GetMonstersByQuery(context.Background(), "equipment.prayer<0")
	// api.GetMonstersByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	GetMonstersByQuery(ctx context.Context, query string) ([]openapi.Monster, error)

	// GetMonstersThatDrop returns a slice of Monsters that drop the given items
	GetMonstersThatDrop(ctx context.Context, items ...string) ([]openapi.Monster, error)

	// GetPrayersByID returns a slice of GetPrayersByID with the given IDs
	GetPrayersByID(ctx context.Context, ids ...string) ([]openapi.Prayer, error)

	// GetPrayersByName returns a slice of GetPrayersByID from the given names
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
	//GetJSONFiles(ctx context.Context, files []string, destinations ...interface{}) error
}

type APIConfig struct {
	UserAgent string
}

var (
	once   sync.Once
	client API
)

// NewAPI creates a osrsboxapi client
// A UserAgent configuration is typically advised
func NewAPI(config APIConfig) API {
	once.Do(func() {
		conf := &openapi.Configuration{
			Scheme:     "https",
			HTTPClient: http.DefaultClient,
			UserAgent:  config.UserAgent,
			Servers: []openapi.ServerConfiguration{
				{
					URL: "api.osrsbox.com",
				},
			},
		}
		client = api.NewAPI(conf)
	})
	return client
}
