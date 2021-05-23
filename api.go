package gosrsbox

import (
	"context"
	"net/http"
	"sync"

	"github.com/atye/gosrsbox/internal/api"
	openapi "github.com/atye/gosrsbox/internal/openapi/api"
	"github.com/atye/gosrsbox/osrsbox"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
)

type API interface {
	// GetItemsByName returns a slice of Items with the given IDs
	GetItemsByID(ctx context.Context, ids ...string) ([]osrsbox.Item, error)

	// GetItemsByName returns a slice of Items from the given wiki names
	GetItemsByName(ctx context.Context, names ...string) ([]osrsbox.Item, error)

	// GetItemsByQuery returns a slice of Items from the given MongoDB or Python query
	// ex:
	// api.GetItemsByQuery(context.Background(), "equipment.prayer<0")
	// api.GetItemsByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	GetItemsByQuery(ctx context.Context, query string) ([]osrsbox.Item, error)

	// GetItemSet returns a slice of Items in the given set
	GetItemSet(ctx context.Context, set sets.SetName) ([]osrsbox.Item, error)

	// GetItemsBySlot returns a slice of Items in the given slot
	GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]osrsbox.Item, error)

	// GetMonstersByID returns a slice of Monsters with the given IDs
	GetMonstersByID(ctx context.Context, ids ...string) ([]osrsbox.Monster, error)

	// GetMonstersByName returns a slice of Monsters from the given wiki names
	GetMonstersByName(ctx context.Context, names ...string) ([]osrsbox.Monster, error)

	// GetMonstersByQuery returns a slice of Monsters from the given MongoDB or Python
	// ex:
	// api.GetMonstersByQuery(context.Background(), "equipment.prayer<0")
	// api.GetMonstersByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	GetMonstersByQuery(ctx context.Context, query string) ([]osrsbox.Monster, error)

	// GetMonstersThatDrop returns a slice of Monsters that drop the given items
	GetMonstersThatDrop(ctx context.Context, items ...string) ([]osrsbox.Monster, error)

	// GetPrayersByID returns a slice of GetPrayersByID with the given IDs
	GetPrayersByID(ctx context.Context, ids ...string) ([]osrsbox.Prayer, error)

	// GetPrayersByName returns a slice of GetPrayersByID from the given names
	GetPrayersByName(ctx context.Context, names ...string) ([]osrsbox.Prayer, error)

	// GetPrayersByQuery returns a slice of Prayers form the given MongoDB or Python query
	// ex:
	// api.GetPrayersByQuery(context.Background(), "equipment.prayer<0")
	// api.GetPrayersByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	GetPrayersByQuery(ctx context.Context, query string) ([]osrsbox.Prayer, error)

	// GetDocument fetches from the Static JSON API and unmarshals the data into the destination.
	// A destination is provided by you for flexibility. The data structures are not quite the same as the responses
	// from the REST API and the REST API doesn't serve everything, such as npcs-summary.json
	// ex:
	// var out map[string]interface{}
	// api.GetDocument(context.Background(), "items-json/0.json", &out)
	// api.GetDocument(context.Background(), "npcs-summary.json", &out)
	GetDocument(ctx context.Context, file string, destination interface{}) error
}

var (
	once   sync.Once
	client API
)

// NewAPI returns a osrsboxapi client.
func NewAPI(userAgent string) API {
	once.Do(func() {
		conf := &openapi.Configuration{
			Scheme:     "https",
			HTTPClient: http.DefaultClient,
			UserAgent:  userAgent,
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
