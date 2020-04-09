package gosrsbox

import (
	"context"
	"fmt"
	"net/http"
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
	// GetMonstersThatDrop returns a slice of monsters that drop the supplied item names
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
	client    HTTPClient
	endpoints *endpoints
}

type endpoints struct {
	items    string
	monsters string
	prayers  string
}

// New returns an OSRSBoxWrapper.
// Pass in nil to use the default http client.
// Pass in an HTTPClient to use your own.
func New(c HTTPClient) OSRSBoxClient {
	if c != nil {
		return &client{
			client: c,
			endpoints: &endpoints{
				items:    items,
				monsters: monsters,
				prayers:  prayers,
			},
		}
	}

	return &client{
		client: &http.Client{},
		endpoints: &endpoints{
			items:    items,
			monsters: monsters,
			prayers:  prayers,
		},
	}
}

func (c *client) GetAllItems(ctx context.Context) ([]*Item, error) {
	itemsI, err := getAll(ctx, c.client, c.endpoints.items)
	if err != nil {
		return nil, err
	}

	if items, ok := itemsI.([]*Item); ok {
		return items, nil
	}

	return nil, fmt.Errorf("Error with type conversion: expected []*gosrsbox.Item but got %T", itemsI)
}

func (c *client) GetItemsByName(ctx context.Context, names ...string) ([]*Item, error) {
	itemsI, err := getByName(ctx, c.client, c.endpoints.items, names...)
	if err != nil {
		return nil, err
	}

	if items, ok := itemsI.([]*Item); ok {
		return items, nil
	}

	return nil, fmt.Errorf("Error with type conversion: expected []*gosrsbox.Item but got %T", itemsI)
}

func (c *client) GetItemsWhere(ctx context.Context, query string) ([]*Item, error) {
	itemsI, err := getWhere(ctx, c.client, c.endpoints.items, query)
	if err != nil {
		return nil, err
	}

	if items, ok := itemsI.([]*Item); ok {
		return items, nil
	}

	return nil, fmt.Errorf("Error with type conversion: expected []*gosrsbox.Item but got %T", itemsI)
}

func (c *client) GetAllMonsters(ctx context.Context) ([]*Monster, error) {
	monstersI, err := getAll(ctx, c.client, c.endpoints.monsters)
	if err != nil {
		return nil, err
	}

	if monsters, ok := monstersI.([]*Monster); ok {
		return monsters, nil
	}

	return nil, fmt.Errorf("Error with type conversion: expected []*gosrsbox.Monster but got %T", monstersI)
}

func (c *client) GetMonstersByName(ctx context.Context, names ...string) ([]*Monster, error) {
	monstersI, err := getByName(ctx, c.client, c.endpoints.monsters, names...)
	if err != nil {
		return nil, err
	}

	if monsters, ok := monstersI.([]*Monster); ok {
		return monsters, nil
	}

	return nil, fmt.Errorf("Error with type conversion: expected []*gosrsbox.Monster but got %T", monstersI)
}

func (c *client) GetMonstersThatDrop(ctx context.Context, names ...string) ([]*Monster, error) {
	monstersI, err := getThatDrop(ctx, c.client, c.endpoints.monsters, names...)
	if err != nil {
		return nil, err
	}

	if monsters, ok := monstersI.([]*Monster); ok {
		return monsters, nil
	}

	return nil, fmt.Errorf("Error with type conversion: expected []*gosrsbox.Monster but got %T", monstersI)
}

func (c *client) GetMonstersWhere(ctx context.Context, query string) ([]*Monster, error) {
	monstersI, err := getWhere(ctx, c.client, c.endpoints.monsters, query)
	if err != nil {
		return nil, err
	}

	if monsters, ok := monstersI.([]*Monster); ok {
		return monsters, nil
	}

	return nil, fmt.Errorf("Error with type conversion: expected []*gosrsbox.Monster but got %T", monstersI)
}

func (c *client) GetAllPrayers(ctx context.Context) ([]*Prayer, error) {
	prayersI, err := getAll(ctx, c.client, c.endpoints.prayers)
	if err != nil {
		return nil, err
	}

	if prayers, ok := prayersI.([]*Prayer); ok {
		return prayers, nil
	}

	return nil, fmt.Errorf("Error with type conversion: expected []*gosrsbox.Prayer but got %T", prayersI)
}

func (c *client) GetPrayersByName(ctx context.Context, names ...string) ([]*Prayer, error) {
	prayersI, err := getByName(ctx, c.client, c.endpoints.prayers, names...)
	if err != nil {
		return nil, err
	}

	if prayers, ok := prayersI.([]*Prayer); ok {
		return prayers, nil
	}

	return nil, fmt.Errorf("Error with type conversion: expected []*gosrsbox.Prayer but got %T", prayersI)
}

func (c *client) GetPrayersWhere(ctx context.Context, query string) ([]*Prayer, error) {
	prayersI, err := getWhere(ctx, c.client, c.endpoints.prayers, query)
	if err != nil {
		return nil, err
	}

	if prayers, ok := prayersI.([]*Prayer); ok {
		return prayers, nil
	}

	return nil, fmt.Errorf("Error with type conversion: expected []*gosrsbox.Prayer but got %T", prayersI)
}
