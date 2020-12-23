package api

import (
	"context"

	"github.com/atye/gosrsbox/osrsboxapi/api/internal/client/openapi"
	"github.com/atye/gosrsbox/osrsboxapi/sets"
	"github.com/atye/gosrsbox/osrsboxapi/slots"
	"golang.org/x/sync/semaphore"
)

type semaphoreMW struct {
	next API
	sem  *semaphore.Weighted
}

func withSemaphore(next API, sem *semaphore.Weighted) semaphoreMW {
	return semaphoreMW{next: next, sem: sem}
}

func (l semaphoreMW) GetItemsByName(ctx context.Context, names ...string) ([]openapi.Item, error) {
	err := l.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer l.sem.Release(1)
	return l.next.GetItemsByName(ctx, names...)
}

func (l semaphoreMW) GetItemsByQuery(ctx context.Context, query string) ([]openapi.Item, error) {
	err := l.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer l.sem.Release(1)
	return l.next.GetItemsByQuery(ctx, query)
}

func (l semaphoreMW) GetItemSet(ctx context.Context, set sets.SetName) ([]openapi.Item, error) {
	err := l.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer l.sem.Release(1)
	return l.next.GetItemSet(ctx, set)
}

func (l semaphoreMW) GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]openapi.Item, error) {
	err := l.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer l.sem.Release(1)
	return l.next.GetItemsBySlot(ctx, slot)
}

func (l semaphoreMW) GetMonstersByName(ctx context.Context, names ...string) ([]openapi.Monster, error) {
	err := l.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer l.sem.Release(1)
	return l.next.GetMonstersByName(ctx, names...)
}

func (l semaphoreMW) GetMonstersByQuery(ctx context.Context, query string) ([]openapi.Monster, error) {
	err := l.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer l.sem.Release(1)
	return l.next.GetMonstersByQuery(ctx, query)
}

func (l semaphoreMW) GetMonstersThatDrop(ctx context.Context, items ...string) ([]openapi.Monster, error) {
	err := l.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer l.sem.Release(1)
	return l.next.GetMonstersThatDrop(ctx, items...)
}

func (l semaphoreMW) GetPrayersByName(ctx context.Context, names ...string) ([]openapi.Prayer, error) {
	err := l.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer l.sem.Release(1)
	return l.next.GetPrayersByName(ctx, names...)
}

func (l semaphoreMW) GetPrayersByQuery(ctx context.Context, query string) ([]openapi.Prayer, error) {
	err := l.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer l.sem.Release(1)
	return l.next.GetPrayersByQuery(ctx, query)
}

func (l semaphoreMW) GetJSONFiles(ctx context.Context, files []string, destinations ...interface{}) error {
	// destinations must be variadic?
	return l.next.GetJSONFiles(ctx, files, destinations...)
}
