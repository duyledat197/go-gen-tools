package tracing

import (
	"context"
	"fmt"
	"io"
	"time"

	cfg "github.com/duyledat197/go-gen-tools/config"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	jeager_zap "github.com/uber/jaeger-client-go/log/zap"
	"go.uber.org/zap"
)

type TracerClient struct {
	ServiceName string
	Address     *cfg.ConnectionAddr

	Tracer    opentracing.Tracer
	config    config.Configuration
	Transport jaeger.Transport
	closer    io.Closer

	Logger *zap.Logger
}

func (c *TracerClient) Connect(ctx context.Context) error {
	cfg := config.Configuration{
		ServiceName: c.ServiceName,
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeRemote,
		},
	}
	transport, err := jaeger.NewUDPTransport(c.Address.GetConnectionString(), 0)
	if err != nil {
		return fmt.Errorf("create jeager transport error: %w", err)
	}
	sampler, err := jaeger.NewGuaranteedThroughputProbabilisticSampler(10, 0.01)
	if err != nil {
		return fmt.Errorf("create jeager sampler error: %w", err)
	}

	options := []config.Option{
		config.Reporter(jaeger.NewRemoteReporter(
			transport,
			jaeger.ReporterOptions.BufferFlushInterval(1*time.Second),
			jaeger.ReporterOptions.Logger(jeager_zap.NewLogger(c.Logger)),
		)),
		config.Sampler(sampler),
	}

	tracer, closer, err := cfg.NewTracer(options...)

	if err != nil {
		return fmt.Errorf("new jeager error: %w", err)
	}

	c.Tracer = tracer
	c.Transport = transport
	c.config = cfg
	c.closer = closer
	return nil
}

func (c *TracerClient) Stop(ctx context.Context) error {
	return c.closer.Close()
}
