package static

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/atye/gosrsbox/osrsboxapi"
	"github.com/atye/gosrsbox/osrsboxapi/api/internal/static/client"
	"github.com/atye/gosrsbox/osrsboxapi/sets"
)

type API interface {
	GetItemsByName(ctx context.Context, names ...string) ([]osrsboxapi.Item, error)
	GetItemsByQuery(ctx context.Context, query string) ([]osrsboxapi.Item, error)
	GetItemSet(ctx context.Context, set sets.SetName) ([]osrsboxapi.Item, error)
	GetMonstersByName(ctx context.Context, names ...string) ([]osrsboxapi.Monster, error)
	GetMonstersByQuery(ctx context.Context, query string) ([]osrsboxapi.Monster, error)
	GetMonstersThatDrop(ctx context.Context, items ...string) ([]osrsboxapi.Monster, error)
	GetPrayersByName(ctx context.Context, names ...string) ([]osrsboxapi.Prayer, error)
	GetPrayersByQuery(ctx context.Context, query string) ([]osrsboxapi.Prayer, error)
	UpdateItems() error
	UpdateMonsters() error
	UpdatePrayers() error
}

type APIConfig struct {
	Logger     *log.Logger
	HttpClient *http.Client
}

func NewAPI(config *APIConfig) (API, error) {
	logger, httpClient := logger(config), httpClient(config)

	api := client.NewAPI(httpClient)
	err := api.RunOptions(client.WithSource(client.FromHttpClient(httpClient)), client.WithOptionLogging(logger, client.WithInit(), "Initializng"))
	if err != nil {
		return nil, err
	}
	return withLogger(api, logger), nil
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
