package gosrsbox

import (
	"context"
	"net/http"
	"sync"
)

// OSRSBoxClient is the osrsbox-api client. Cancellation should be controlled through the passed in context.
type OSRSBoxClient interface {
	// GetAllItems returns a slice of all items.
	GetAllItems(ctx context.Context) ([]*Item, error)
	// GetItemsByName returns a slice of items specified by name.
	GetItemsByName(ctx context.Context, names ...string) ([]*Item, error)
	// GetItemsWhere returns a slice of items from the supplied MongoDB query.
	GetItemsWhere(ctx context.Context, query string) ([]*Item, error)
	// GetAllMonsters returns a slice of all monsters.
	GetAllMonsters(ctx context.Context) ([]*Monster, error)
	// GetMonstersByName returns a slice of monsters specified by name.
	GetMonstersByName(ctx context.Context, names ...string) ([]*Monster, error)
	// GetMonstersThatDrop returns a slice of monsters that drop the supplied item names.
	// It returns monsters that drop any of the item names, not monsters that drop all of them.
	GetMonstersThatDrop(ctx context.Context, names ...string) ([]*Monster, error)
	// GetMonstersWhere returns a slice of monsters from the supplied MongoDB query.
	GetMonstersWhere(ctx context.Context, query string) ([]*Monster, error)
	// GetAllPrayers returns a slice of all prayers.
	GetAllPrayers(ctx context.Context) ([]*Prayer, error)
	// GetPrayersByName returns a slice of prayers specified by name.
	GetPrayersByName(ctx context.Context, names ...string) ([]*Prayer, error)
	// GetPrayersWhere returns a slice of prayers from the supplied MongoDB query.
	GetPrayersWhere(ctx context.Context, query string) ([]*Prayer, error)
}

// HTTPClient is the client for doing HTTP calls.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type client struct {
	client HTTPClient
	wg     sync.WaitGroup
	mu     sync.Mutex
}

// New returns an OSRSBoxClient.
// Pass in nil to use the default http client.
// Pass in an HTTPClient to use your own.
func New(c HTTPClient) OSRSBoxClient {
	if c != nil {
		return &client{
			client: c,
		}
	}

	return &client{
		client: &http.Client{},
	}
}

func (c *client) GetAllItems(ctx context.Context) ([]*Item, error) {
	return getAllItems(ctx, c)
}

func (c *client) GetItemsByName(ctx context.Context, names ...string) ([]*Item, error) {
	return getItemsByName(ctx, c, names...)
}

func (c *client) GetItemsWhere(ctx context.Context, query string) ([]*Item, error) {
	return getItemsWhere(ctx, c, query)
}

func (c *client) GetAllMonsters(ctx context.Context) ([]*Monster, error) {
	return getAllMonsters(ctx, c)
}

func (c *client) GetMonstersByName(ctx context.Context, names ...string) ([]*Monster, error) {
	return getMonstersByName(ctx, c, names...)
}

func (c *client) GetMonstersThatDrop(ctx context.Context, names ...string) ([]*Monster, error) {
	return getMonstersThatDrop(ctx, c, names...)
}

func (c *client) GetMonstersWhere(ctx context.Context, query string) ([]*Monster, error) {
	return getMonstersWhere(ctx, c, query)
}

func (c *client) GetAllPrayers(ctx context.Context) ([]*Prayer, error) {
	return getAllPrayers(ctx, c.client)
}

func (c *client) GetPrayersByName(ctx context.Context, names ...string) ([]*Prayer, error) {
	return getPrayersByName(ctx, c.client, names...)
}

func (c *client) GetPrayersWhere(ctx context.Context, query string) ([]*Prayer, error) {
	return getPrayersWhere(ctx, c.client, query)
}
