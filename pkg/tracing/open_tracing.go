package tracing

import (
	"fmt"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type OpenTracer struct {
	ServiceName string
	Address     string
	Tracer      opentracing.Tracer
}

func NewOpenTracer(serviceName, address string) (*OpenTracer, error) {
	cfg := config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeRemote,
		},
	}
	sender, err := jaeger.NewUDPTransport(address, 0)
	if err != nil {
		return nil, err
	}
	sampler, _ := jaeger.NewGuaranteedThroughputProbabilisticSampler(10, 0.01)

	options := []config.Option{
		config.Reporter(jaeger.NewRemoteReporter(
			sender,
			jaeger.ReporterOptions.BufferFlushInterval(1*time.Second),
		)),
		config.Sampler(sampler),
	}

	tracer, _, err := cfg.NewTracer(options...)

	if err != nil {
		return nil, fmt.Errorf("tracer: %w", err)
	}

	return &OpenTracer{
		ServiceName: serviceName,
		Address:     address,
		Tracer:      tracer,
	}, nil
}
