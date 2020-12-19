package osrsbox

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/atye/gosrsbox/osrsbox/db"
	"github.com/atye/gosrsbox/osrsbox/db/backends/external"
	externalMW "github.com/atye/gosrsbox/osrsbox/db/backends/external/middleware"
	"github.com/atye/gosrsbox/osrsbox/db/sets"
)

type ItemGetter interface {
	GetItemsByName(ctx context.Context, names ...string) ([]db.Item, error)
	GetItemsByQuery(ctx context.Context, query string) ([]db.Item, error)
	GetItemSet(ctx context.Context, set sets.SetName) ([]db.Item, error)
}

type MonsterGetter interface {
	GetMonstersByName(ctx context.Context, names ...string) ([]db.Monster, error)
	GetMonstersByQuery(ctx context.Context, query string) ([]db.Monster, error)
	GetMonstersThatDrop(ctx context.Context, items ...string) ([]db.Monster, error)
}

type PrayerGetter interface {
	GetPrayersByName(ctx context.Context, names ...string) ([]db.Prayer, error)
	GetPrayersByQuery(ctx context.Context, query string) ([]db.Prayer, error)
}

type Updater interface {
	UpdateItems() error
	UpdateMonsters() error
	UpdatePrayers() error
}

type InMemoryAPI interface {
	ItemGetter
	MonsterGetter
	PrayerGetter
	Updater
}

type InMemoryAPIConfig struct {
	Logger *log.Logger
}

type ExternalAPI interface {
	ItemGetter
	MonsterGetter
	PrayerGetter
}

type ExternalAPIConfig struct {
	Logger *log.Logger
}

/*func NewInMemoryAPI(config *InMemoryAPIConfig) (InMemoryAPI, error) {
	var logger *log.Logger
	if config == nil {
		logger = log.New(os.Stdout, "osrsbox", log.LstdFlags)
	} else if config.Logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	} else {
		logger = config.Logger
	}

	api := inmemory.NewAPI()
	err := api.RunOptions(inmemory.WithSource(inmemory.FromRawGithub()), inmemory.WithOptionLogging(logger, inmemory.WithInit()))
	if err != nil {
		return nil, err
	}
	return inmemoryMW.LoggingMW(api, logger), nil
}*/

var (
	once sync.Once
	api  ExternalAPI
)

func NewAPI(config *ExternalAPIConfig) ExternalAPI {
	once.Do(func() {
		var logger *log.Logger
		if config == nil {
			logger = log.New(os.Stdout, "osrsbox", log.LstdFlags)
		} else if config.Logger == nil {
			logger = log.New(ioutil.Discard, "", log.LstdFlags)
		} else {
			logger = config.Logger
		}
		api = externalMW.LoggingMW(external.NewAPI(), logger)
	})
	return api
}
