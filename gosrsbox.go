package gosrsbox

import (
	"io"
	"log"
	"sync"

	osrsboxapi "github.com/atye/gosrsbox/api"
	"github.com/atye/gosrsbox/internal/client"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
)

var (
	once      sync.Once
	apiClient osrsboxapi.API
)

// Option to configure the client
type Option func(c *client.APIClient) *client.APIClient

// WithTracing enables Zipkin tracing.
// Pass in the collector URL and a probablity to export traces
func WithTracing(url string, probablity float64) Option {
	return func(c *client.APIClient) *client.APIClient {
		err := initTracing(url, probablity)
		if err != nil {
			panic(err)
		}
		return c
	}
}

// NewAPI returns a osrsboxapi client.
func NewAPI(userAgent string, opts ...Option) osrsboxapi.API {
	once.Do(func() {
		c := client.NewAPI(userAgent)
		for _, o := range opts {
			c = o(c)
		}
		apiClient = c
	})
	return apiClient
}

func initTracing(url string, probability float64) error {
	exporter, err := zipkin.NewRawExporter(
		url,
		zipkin.WithLogger(log.New(io.Discard, "", log.LstdFlags)),
		zipkin.WithSDKOptions(trace.WithSampler(trace.TraceIDRatioBased(probability))),
	)
	if err != nil {
		return err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultBatchTimeout),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
		),
		trace.WithResource(resource.NewWithAttributes(
			semconv.ServiceNameKey.String("gosrsbox"),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
