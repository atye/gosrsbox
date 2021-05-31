package common

import (
	"context"

	"github.com/atye/gosrsbox/models"
)

type Params struct {
	Where string
	Page  int
}

type RequestExecutor interface {
	ExecuteItemsRequest(context.Context, Params) (ItemsResponse, error)
	ExecuteMonstersRequest(context.Context, Params) (MonstersResponse, error)
	ExecutePrayersRequest(context.Context, Params) (PrayersResponse, error)
}

type MetaGetter interface {
	GetTotal() int
	GetMaxResults() int
}

type ItemsResponse interface {
	MetaGetter
	GetItems() []models.Item
}

type MonstersResponse interface {
	MetaGetter
	GetMonsters() []models.Monster
}

type PrayersResponse interface {
	MetaGetter
	GetPrayers() []models.Prayer
}
