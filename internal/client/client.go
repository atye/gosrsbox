package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/atye/gosrsbox/api"
	"github.com/atye/gosrsbox/internal/common"
	"github.com/atye/gosrsbox/internal/openapi"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/semaphore"
)

type APIClient struct {
	docsAddress string
	reqExecutor common.RequestExecutor
	sem         *semaphore.Weighted
	tracer      trace.Tracer
}

var _ api.API = &APIClient{}

const (
	jsonDocuments = "https://www.osrsbox.com/osrsbox-db/"
	osrsboxpi     = "api.osrsbox.com"
)

var (
	errNoIDs   = errors.New("no ids provided")
	errNoNames = errors.New("no names provided")
	errNoSet   = errors.New("no set provided")
	errNoSlot  = errors.New("no slot provided")
)

func NewAPI(userAgent string) *APIClient {
	return &APIClient{
		docsAddress: jsonDocuments,
		reqExecutor: openapi.NewClient(userAgent, "https", osrsboxpi),
		sem:         semaphore.NewWeighted(int64(10)),
		tracer:      otel.GetTracerProvider().Tracer("gosrsbox"),
	}
}

func (c *APIClient) doItemsRequest(ctx context.Context, p common.Params) (common.ItemsResponse, error) {
	ctx, span := c.createSpan(ctx, "execute_items_request")
	defer span.End()

	setSpanAttributesFromParams(span, p)

	err := c.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer c.sem.Release(1)

	resp, err := c.reqExecutor.ExecuteItemsRequest(ctx, p)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *APIClient) doMonstersRequest(ctx context.Context, p common.Params) (common.MonstersResponse, error) {
	ctx, span := c.createSpan(ctx, "execute_monsters_request")
	defer span.End()

	setSpanAttributesFromParams(span, p)

	err := c.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer c.sem.Release(1)

	resp, err := c.reqExecutor.ExecuteMonstersRequest(ctx, p)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *APIClient) doPrayersRequest(ctx context.Context, p common.Params) (common.PrayersResponse, error) {
	ctx, span := c.createSpan(ctx, "execute_prayers_request")
	defer span.End()

	setSpanAttributesFromParams(span, p)

	err := c.sem.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}
	defer c.sem.Release(1)

	resp, err := c.reqExecutor.ExecutePrayersRequest(ctx, p)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *APIClient) doDocumentRequest(ctx context.Context, url string) (*http.Response, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "execute_document_request")
	defer span.End()

	span.SetAttributes(attribute.String("url", url))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		return nil, fmt.Errorf("code: %d, message: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return resp, nil
}

func quoteStrings(elements ...string) []string {
	quotedStrings := make([]string, len(elements))
	for i, e := range elements {
		quotedStrings[i] = fmt.Sprintf("%q", e)
	}
	return quotedStrings
}

func (c *APIClient) createSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	if trace.SpanFromContext(ctx).SpanContext().HasSpanID() {
		return trace.SpanFromContext(ctx).Tracer().Start(ctx, name)
	}
	return c.tracer.Start(ctx, name)
}

func setSpanAttributesFromParams(span trace.Span, p common.Params) {
	if p.Where != "" {
		span.SetAttributes(attribute.String("where", p.Where))
	}
	if p.Page != 0 {
		span.SetAttributes(attribute.Int("page", p.Page))
	}
}

func setSpanErrorStatus(span trace.Span, err error) {
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
}
