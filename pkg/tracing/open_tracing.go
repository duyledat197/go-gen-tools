package tracing

import (
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

type OpenTracer struct {
	ServiceName string
	Address     string

	Tracer    opentracing.Tracer
	Config    config.Configuration
	Transport jaeger.Transport

	Logger *zap.Logger
}

func (t *OpenTracer) Init() *OpenTracer {
	cfg := config.Configuration{
		ServiceName: t.ServiceName,
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeRemote,
		},
	}
	transport, err := jaeger.NewUDPTransport(t.Address, 0)
	if err != nil {
		t.Logger.Panic("create jeager transport error: ", zap.Error(err))
	}
	sampler, err := jaeger.NewGuaranteedThroughputProbabilisticSampler(10, 0.01)
	if err != nil {
		t.Logger.Panic("create jeager sampler error: ", zap.Error(err))
	}

	options := []config.Option{
		config.Reporter(jaeger.NewRemoteReporter(
			transport,
			jaeger.ReporterOptions.BufferFlushInterval(1*time.Second),
		)),
		config.Sampler(sampler),
	}

	tracer, _, err := cfg.NewTracer(options...)

	if err != nil {
		if err != nil {
			t.Logger.Panic("new jeager error: ", zap.Error(err))
		}
	}

	t.Tracer = tracer
	t.Transport = transport
	t.Config = cfg

	return t
}
