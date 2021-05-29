package client

import (
	"context"

	osrsboxapi "github.com/atye/gosrsbox/api"
	"github.com/atye/gosrsbox/models"
	"github.com/atye/gosrsbox/sets"
	"github.com/atye/gosrsbox/slots"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type tracingMW struct {
	c      osrsboxapi.API
	tracer trace.Tracer
}

func NewTracingMW(c osrsboxapi.API) osrsboxapi.API {
	return &tracingMW{
		c:      c,
		tracer: otel.GetTracerProvider().Tracer("gosrsbox"),
	}
}

func (mw *tracingMW) GetItemsByID(ctx context.Context, ids ...string) ([]models.Item, error) {
	ctx, span := mw.tracer.Start(ctx, "get_items_by_id")
	defer span.End()
	return mw.c.GetItemsByID(ctx, ids...)
}

func (mw *tracingMW) GetItemsByName(ctx context.Context, names ...string) ([]models.Item, error) {
	ctx, span := mw.tracer.Start(ctx, "get_items_by_name")
	defer span.End()
	return mw.c.GetItemsByName(ctx, names...)
}

func (mw *tracingMW) GetItemSet(ctx context.Context, set sets.SetName) ([]models.Item, error) {
	ctx, span := mw.tracer.Start(ctx, "get_item_set")
	defer span.End()
	return mw.c.GetItemSet(ctx, set)
}

func (mw *tracingMW) GetItemsBySlot(ctx context.Context, slot slots.SlotName) ([]models.Item, error) {
	ctx, span := mw.tracer.Start(ctx, "get_items_by_slot")
	defer span.End()
	return mw.c.GetItemsBySlot(ctx, slot)
}

func (mw *tracingMW) GetItemsByQuery(ctx context.Context, query string) ([]models.Item, error) {
	ctx, span := mw.tracer.Start(ctx, "get_items_by_query")
	defer span.End()
	return mw.c.GetItemsByQuery(ctx, query)
}

func (mw *tracingMW) GetMonstersByID(ctx context.Context, ids ...string) ([]models.Monster, error) {
	ctx, span := mw.tracer.Start(ctx, "get_monsters_by_id")
	defer span.End()
	return mw.c.GetMonstersByID(ctx, ids...)
}

func (mw *tracingMW) GetMonstersByName(ctx context.Context, names ...string) ([]models.Monster, error) {
	ctx, span := mw.tracer.Start(ctx, "get_monsters_by_name")
	defer span.End()
	return mw.c.GetMonstersByName(ctx, names...)
}

func (mw *tracingMW) GetMonstersThatDrop(ctx context.Context, names ...string) ([]models.Monster, error) {
	ctx, span := mw.tracer.Start(ctx, "get_monsters_that_drop")
	defer span.End()
	return mw.c.GetMonstersThatDrop(ctx, names...)
}

func (mw *tracingMW) GetMonstersByQuery(ctx context.Context, query string) ([]models.Monster, error) {
	ctx, span := mw.tracer.Start(ctx, "get_monsters_by_query")
	defer span.End()
	return mw.c.GetMonstersByQuery(ctx, query)
}

func (mw *tracingMW) GetPrayersByID(ctx context.Context, ids ...string) ([]models.Prayer, error) {
	ctx, span := mw.tracer.Start(ctx, "get_prayers_by_id")
	defer span.End()
	return mw.c.GetPrayersByID(ctx, ids...)
}

func (mw *tracingMW) GetPrayersByName(ctx context.Context, names ...string) ([]models.Prayer, error) {
	ctx, span := mw.tracer.Start(ctx, "get_prayers_by_name")
	defer span.End()
	return mw.c.GetPrayersByName(ctx, names...)
}

func (mw *tracingMW) GetPrayersByQuery(ctx context.Context, query string) ([]models.Prayer, error) {
	ctx, span := mw.tracer.Start(ctx, "get_prayers_by_query")
	defer span.End()
	return mw.c.GetPrayersByQuery(ctx, query)
}

func (mw *tracingMW) GetDocument(ctx context.Context, file string, destination interface{}) error {
	ctx, span := mw.tracer.Start(ctx, "get_document")
	defer span.End()
	return mw.c.GetDocument(ctx, file, destination)
}
