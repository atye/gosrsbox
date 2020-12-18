package osrsbox

import (
	"context"
	"net/http"

	"github.com/atye/gosrsbox/osrsbox/db"
	"github.com/atye/gosrsbox/osrsbox/db/api"
	inmemory "github.com/atye/gosrsbox/osrsbox/db/in-memory"
	"github.com/atye/gosrsbox/osrsbox/db/update"
)

type itemGetter interface {
	GetItemsByName(ctx context.Context, names ...string) ([]db.Item, error)
	GetItemsByQuery(ctx context.Context, query string) ([]db.Item, error)
}

type monsterGetter interface {
	GetMonstersByName(ctx context.Context, names ...string) ([]db.Monster, error)
	GetMonstersByQuery(ctx context.Context, query string) ([]db.Monster, error)
	GetMonstersThatDrop(ctx context.Context, items ...string) ([]db.Monster, error)
}

type prayerGetter interface {
	GetPrayersByName(ctx context.Context, names ...string) ([]db.Prayer, error)
	GetPrayersByQuery(ctx context.Context, query string) ([]db.Prayer, error)
}

type updater interface {
	UpdateItems() error
	UpdateMonsters() error
	UpdatePrayers() error
}

type InMemoryClient interface {
	itemGetter
	monsterGetter
	prayerGetter
	updater
}

type APIClient interface {
	itemGetter
	monsterGetter
	prayerGetter
}

func NewInMemoryClient() (InMemoryClient, error) {
	api, err := inmemory.NewInMemoryClient(inmemory.WithUpdater(update.HttpGet(http.Get)), inmemory.WithInit())
	if err != nil {
		return nil, err
	}
	return api, nil
}

func NewAPIClient() APIClient {
	return api.NewAPIClient()
}
