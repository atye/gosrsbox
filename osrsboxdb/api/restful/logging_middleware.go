package restful

import (
	"context"
	"log"
	"time"

	"github.com/atye/gosrsbox/osrsboxdb"
	"github.com/atye/gosrsbox/osrsboxdb/sets"
)

type loggingMW struct {
	next API
	log  *log.Logger
}

func withLogger(next API, log *log.Logger) loggingMW {
	return loggingMW{next: next, log: log}
}

func (l loggingMW) GetItemsByName(ctx context.Context, names ...string) ([]osrsboxdb.Item, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetItemsByName took %v", time.Since(now))
	}()
	return l.next.GetItemsByName(ctx, names...)
}

func (l loggingMW) GetItemsByQuery(ctx context.Context, query string) ([]osrsboxdb.Item, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetItemsByQuery took %v", time.Since(now))
	}()
	return l.next.GetItemsByQuery(ctx, query)
}

func (l loggingMW) GetItemSet(ctx context.Context, set sets.SetName) ([]osrsboxdb.Item, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetItemSet took %v", time.Since(now))
	}()
	return l.next.GetItemSet(ctx, set)
}

func (l loggingMW) GetMonstersByName(ctx context.Context, names ...string) ([]osrsboxdb.Monster, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetMonstersByName took %v", time.Since(now))
	}()
	return l.next.GetMonstersByName(ctx, names...)
}

func (l loggingMW) GetMonstersByQuery(ctx context.Context, query string) ([]osrsboxdb.Monster, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetMonstersByQuery took %v", time.Since(now))
	}()
	return l.next.GetMonstersByQuery(ctx, query)
}

func (l loggingMW) GetMonstersThatDrop(ctx context.Context, items ...string) ([]osrsboxdb.Monster, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetMonstersThatDrop took %v", time.Since(now))
	}()
	return l.next.GetMonstersThatDrop(ctx, items...)
}

func (l loggingMW) GetPrayersByName(ctx context.Context, names ...string) ([]osrsboxdb.Prayer, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetPrayersByName took %v", time.Since(now))
	}()
	return l.next.GetPrayersByName(ctx, names...)
}

func (l loggingMW) GetPrayersByQuery(ctx context.Context, query string) ([]osrsboxdb.Prayer, error) {
	now := time.Now()
	defer func() {
		l.log.Printf("GetPrayersByQuery took %v", time.Since(now))
	}()
	return l.next.GetPrayersByQuery(ctx, query)
}
