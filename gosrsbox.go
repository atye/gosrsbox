package gosrsbox

import (
	"io/ioutil"
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

type Option func(c *client.APIClient) *client.APIClient

func WithTracing(url string) Option {
	return func(c *client.APIClient) *client.APIClient {
		err := initTracing(url)
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

func initTracing(url string) error {
	exporter, err := zipkin.NewRawExporter(
		url,
		zipkin.WithLogger(log.New(ioutil.Discard, "", log.LstdFlags)),
		zipkin.WithSDKOptions(trace.WithSampler(trace.AlwaysSample())),
	)

	if err != nil {
		return err
	}

	bsp := trace.NewBatchSpanProcessor(exporter)

	tp := trace.NewTracerProvider(
		trace.WithSpanProcessor(bsp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.ServiceNameKey.String("gosrsbox"),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
