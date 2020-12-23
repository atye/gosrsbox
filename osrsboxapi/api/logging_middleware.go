package api

import (
	"context"
	"log"
	"time"

	openapi "github.com/atye/gosrsbox/osrsboxapi/openapi/api"
	"github.com/atye/gosrsbox/osrsboxapi/sets"
	"github.com/atye/gosrsbox/osrsboxapi/slots"
)

type loggingMW struct {
	next API
	log  *log.Logger
}

func withLogger(next API, log *log.Logger) loggingMW {
	return loggingMW{next: next, log: log}
}

func (l loggingMW) GetItemsByID(ctx context.Context, ids ...string) ([]openapi.Item, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetItemByID took %v", time.Since(now))
	}()
	return l.next.GetItemsByID(ctx, ids...)
}

func (l loggingMW) GetItemsByName(ctx context.Context, names ...string) ([]openapi.Item, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetItemsByName took %v", time.Since(now))
	}()
	return l.next.GetItemsByName(ctx, names...)
}

func (l loggingMW) GetItemsByQuery(ctx context.Context, query string) ([]openapi.Item, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetItemsByQuery took %v", time.Since(now))
	}()
	return l.next.GetItemsByQuery(ctx, query)
}

func (l loggingMW) GetItemSet(ctx context.Context, set sets.SetName) ([]openapi.Item, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetItemSet took %v", time.Since(now))
	}()
	return l.next.GetItemSet(ctx, set)
}

func (l loggingMW) GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]openapi.Item, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetItemsBySlot took %v", time.Since(now))
	}()
	return l.next.GetItemsBySlot(ctx, slot)
}

func (l loggingMW) GetMonstersByID(ctx context.Context, ids ...string) ([]openapi.Monster, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetMonstersByID took %v", time.Since(now))
	}()
	return l.next.GetMonstersByID(ctx, ids...)
}

func (l loggingMW) GetMonstersByName(ctx context.Context, names ...string) ([]openapi.Monster, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetMonstersByName took %v", time.Since(now))
	}()
	return l.next.GetMonstersByName(ctx, names...)
}

func (l loggingMW) GetMonstersByQuery(ctx context.Context, query string) ([]openapi.Monster, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetMonstersByQuery took %v", time.Since(now))
	}()
	return l.next.GetMonstersByQuery(ctx, query)
}

func (l loggingMW) GetMonstersThatDrop(ctx context.Context, items ...string) ([]openapi.Monster, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetMonstersThatDrop took %v", time.Since(now))
	}()
	return l.next.GetMonstersThatDrop(ctx, items...)
}

func (l loggingMW) GetPrayersByID(ctx context.Context, ids ...string) ([]openapi.Prayer, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetPrayersByID took %v", time.Since(now))
	}()
	return l.next.GetPrayersByID(ctx, ids...)
}

func (l loggingMW) GetPrayersByName(ctx context.Context, names ...string) ([]openapi.Prayer, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetPrayersByName took %v", time.Since(now))
	}()
	return l.next.GetPrayersByName(ctx, names...)
}

func (l loggingMW) GetPrayersByQuery(ctx context.Context, query string) ([]openapi.Prayer, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetPrayersByQuery took %v", time.Since(now))
	}()
	return l.next.GetPrayersByQuery(ctx, query)
}
