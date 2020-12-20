package restful

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/atye/gosrsbox/osrsboxdb"
	"github.com/atye/gosrsbox/osrsboxdb/api/restful/internal/client"
	"github.com/atye/gosrsbox/osrsboxdb/sets"
)

type API interface {
	GetItemsByName(ctx context.Context, names ...string) ([]osrsboxdb.Item, error)
	GetItemsByQuery(ctx context.Context, query string) ([]osrsboxdb.Item, error)
	GetItemSet(ctx context.Context, set sets.SetName) ([]osrsboxdb.Item, error)
	GetMonstersByName(ctx context.Context, names ...string) ([]osrsboxdb.Monster, error)
	GetMonstersByQuery(ctx context.Context, query string) ([]osrsboxdb.Monster, error)
	GetMonstersThatDrop(ctx context.Context, items ...string) ([]osrsboxdb.Monster, error)
	GetPrayersByName(ctx context.Context, names ...string) ([]osrsboxdb.Prayer, error)
	GetPrayersByQuery(ctx context.Context, query string) ([]osrsboxdb.Prayer, error)
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
		api = withLogger(client.NewAPI(httpClient), logger)
	})
	return api
}

func logger(c *APIConfig) *log.Logger {
	var logger *log.Logger
	if c == nil {
		logger = log.New(os.Stdout, "osrsbox", log.LstdFlags)
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
