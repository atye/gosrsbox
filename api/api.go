package api

import (
	"context"

	"github.com/atye/gosrsbox/models"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
)

type API interface {
	// GetItemsByName returns a slice of Items with the given IDs
	GetItemsByID(ctx context.Context, ids ...string) ([]models.Item, error)

	// GetItemsByName returns a slice of Items from the given wiki names
	GetItemsByName(ctx context.Context, names ...string) ([]models.Item, error)

	// GetItemsByQuery returns a slice of Items from the given MongoDB or Python query
	// ex:
	// api.GetItemsByQuery(context.Background(), "equipment.prayer<0")
	// api.GetItemsByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	GetItemsByQuery(ctx context.Context, query string) ([]models.Item, error)

	// GetItemSet returns a slice of Items in the given set
	GetItemSet(ctx context.Context, set sets.SetName) ([]models.Item, error)

	// GetItemsBySlot returns a slice of Items in the given slot
	GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]models.Item, error)

	// GetMonstersByID returns a slice of Monsters with the given IDs
	GetMonstersByID(ctx context.Context, ids ...string) ([]models.Monster, error)

	// GetMonstersByName returns a slice of Monsters from the given wiki names
	GetMonstersByName(ctx context.Context, names ...string) ([]models.Monster, error)

	// GetMonstersByQuery returns a slice of Monsters from the given MongoDB or Python
	// ex:
	// api.GetMonstersByQuery(context.Background(), "equipment.prayer<0")
	// api.GetMonstersByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	GetMonstersByQuery(ctx context.Context, query string) ([]models.Monster, error)

	// GetMonstersThatDrop returns a slice of Monsters that drop the given items
	GetMonstersThatDrop(ctx context.Context, items ...string) ([]models.Monster, error)

	// GetPrayersByID returns a slice of GetPrayersByID with the given IDs
	GetPrayersByID(ctx context.Context, ids ...string) ([]models.Prayer, error)

	// GetPrayersByName returns a slice of GetPrayersByID from the given names
	GetPrayersByName(ctx context.Context, names ...string) ([]models.Prayer, error)

	// GetPrayersByQuery returns a slice of Prayers form the given MongoDB or Python query
	// ex:
	// api.GetPrayersByQuery(context.Background(), "equipment.prayer<0")
	// api.GetPrayersByQuery(context.Background(), `{ "equipment.prayer": { "$lt": 0 }, "duplicate": false }`)
	GetPrayersByQuery(ctx context.Context, query string) ([]models.Prayer, error)

	// GetDocument fetches the file from the Static JSON API and unmarshals the data into the destination.
	// A destination (should be a pointer) is provided by you for flexibility. The data structures are not quite the same as the responses
	// from the REST API and the REST API doesn't serve everything, such as npcs-summary.json.
	// This is useful for dumping large contents of the database.
	// ex:
	// var out map[string]interface{}
	// api.GetDocument(context.Background(), "items-json/0.json", &out)
	// api.GetDocument(context.Background(), "npcs-summary.json", &out)
	GetDocument(ctx context.Context, file string, destination interface{}) error
}
