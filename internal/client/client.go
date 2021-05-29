package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	osrsboxapi "github.com/atye/gosrsbox/api"
	"github.com/atye/gosrsbox/internal/api"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/semaphore"
)

type apiClient struct {
	docsAddress   string
	openAPIClient *api.APIClient
	sem           *semaphore.Weighted
}

const (
	jsonDocuments = "https://www.osrsbox.com/osrsbox-db/"
)

var (
	errNoIDs   = errors.New("no ids provided")
	errNoNames = errors.New("no names provided")
	errNoSet   = errors.New("no set provided")
	errNoSlot  = errors.New("no slot provided")
)

func NewAPI(conf *api.Configuration) osrsboxapi.API {
	return &apiClient{
		docsAddress:   jsonDocuments,
		openAPIClient: api.NewAPIClient(conf),
		sem:           semaphore.NewWeighted(int64(10)),
	}
}

type params struct {
	where string
	page  int
}

func (c *apiClient) executeItemsRequest(ctx context.Context, p params) (api.InlineResponse200, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "execute_items_request")
	defer span.End()

	if p.where != "" {
		span.SetAttributes(attribute.String("where", p.where))
	}
	if p.page != 0 {
		span.SetAttributes(attribute.Int("page", p.page))
	}

	err := c.sem.Acquire(ctx, 1)
	if err != nil {
		return api.InlineResponse200{}, err
	}
	defer c.sem.Release(1)

	r := c.openAPIClient.ItemApi.Getitems(ctx).Where(p.where)
	if p.page != 0 {
		r = r.Page(int32(p.page))
	}

	inline, resp, err := r.Execute()
	err = checkError(err)
	if err != nil {
		return api.InlineResponse200{}, err
	}
	defer resp.Body.Close()

	return inline, nil
}

func (c *apiClient) doMonstersRequest(ctx context.Context, p params) (api.InlineResponse2003, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "execute_monsters_request")
	defer span.End()

	if p.where != "" {
		span.SetAttributes(attribute.String("where", p.where))
	}
	if p.page != 0 {
		span.SetAttributes(attribute.Int("page", p.page))
	}

	err := c.sem.Acquire(ctx, 1)
	if err != nil {
		return api.InlineResponse2003{}, err
	}
	defer c.sem.Release(1)

	r := c.openAPIClient.MonsterApi.Getmonsters(ctx).Where(p.where)
	if p.page != 0 {
		r = r.Page(int32(p.page))
	}

	inline, resp, err := r.Execute()
	err = checkError(err)
	if err != nil {
		return api.InlineResponse2003{}, err
	}
	defer resp.Body.Close()

	return inline, nil
}

func (c *apiClient) doPrayersRequest(ctx context.Context, p params) (api.InlineResponse2004, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "execute_prayers_request")
	defer span.End()

	if p.where != "" {
		span.SetAttributes(attribute.String("where", p.where))
	}
	if p.page != 0 {
		span.SetAttributes(attribute.Int("page", p.page))
	}

	err := c.sem.Acquire(ctx, 1)
	if err != nil {
		return api.InlineResponse2004{}, err
	}
	defer c.sem.Release(1)

	r := c.openAPIClient.PrayerApi.Getprayers(ctx).Where(p.where)
	if p.page != 0 {
		r = r.Page(int32(p.page))
	}

	inline, resp, err := r.Execute()
	err = checkError(err)
	if err != nil {
		return api.InlineResponse2004{}, err
	}
	defer resp.Body.Close()

	return inline, nil
}

func (c *apiClient) doDocumentRequest(ctx context.Context, url string) (*http.Response, error) {
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

func checkError(executeErr error) error {
	if executeErr == nil {
		return nil
	}

	var genericErr api.GenericOpenAPIError
	if !errors.As(executeErr, &genericErr) {
		return executeErr
	}

	var apiErr api.Error
	err := json.Unmarshal(genericErr.Body(), &apiErr)
	if err != nil {
		return err
	}

	if apiErr.Error.GetCode() == 0 && apiErr.Error.GetMessage() == "" {
		return executeErr
	}
	return fmt.Errorf("code %d, message: %s", apiErr.Error.GetCode(), apiErr.Error.GetMessage())
}

func quoteStrings(elements ...string) []string {
	quotedStrings := make([]string, len(elements))
	for i, e := range elements {
		quotedStrings[i] = fmt.Sprintf("%q", e)
	}
	return quotedStrings
}
