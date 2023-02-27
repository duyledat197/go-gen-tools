package grpc_client

import (
	"context"
	"time"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/duyledat197/go-gen-tools/pkg/registry"
	"github.com/duyledat197/go-gen-tools/pkg/tracing"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
)

type Options struct {
	//* for enable netflix hystrix (circuitbreaker)
	IsEnableHystrix bool

	//* for enable client loadbalancer consul
	IsEnableClientLoadBalancer bool

	//* for enable open tracing
	IsEnableTracing bool

	//* enable retry mode
	IsEnableRetry bool

	//* enable metrics with prometheus
	IsEnableMetrics bool

	//* enable tls secure
	IsEnableSecure bool

	//* for enable validator
	IsEnableValidator bool
}

type GrpcClient struct {
	ServiceName  string
	Tracer       *tracing.TracerClient
	Consul       *registry.ConsulClient
	Address      *config.ConnectionAddr
	Endpoint     string
	Options      *Options
	Creds        credentials.TransportCredentials
	OtherOptions []grpc.DialOption
	Client       *grpc.ClientConn
}

func (c *GrpcClient) Connect(ctx context.Context) error {
	optsRetry := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(50 * time.Millisecond)),
		grpc_retry.WithCodes(codes.Unavailable),
		grpc_retry.WithMax(3),
		grpc_retry.WithPerRetryTimeout(3 * time.Second),
	}
	var (
		streamInterceptors []grpc.StreamClientInterceptor
		unaryInterceptors  []grpc.UnaryClientInterceptor
		opts               []grpc.DialOption
	)

	if c.Options != nil {
		options := c.Options
		if options.IsEnableClientLoadBalancer {
			opts = append(opts, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
			c.Endpoint = c.Consul.GetURL(c.ServiceName)
		} else {
			c.Endpoint = c.Address.GetConnectionString()
		}

		if options.IsEnableHystrix {
			unaryInterceptors = append(unaryInterceptors,
				UnaryClientInterceptor(c.Options.IsEnableHystrix))
		}

		if options.IsEnableTracing {
			streamInterceptors = append(streamInterceptors,
				grpc_opentracing.StreamClientInterceptor(
					grpc_opentracing.WithTracer(c.Tracer.Tracer),
				),
			)

			unaryInterceptors = append(unaryInterceptors,
				grpc_opentracing.UnaryClientInterceptor(
					grpc_opentracing.WithTracer(c.Tracer.Tracer),
				),
			)
		}

		if options.IsEnableRetry {
			streamInterceptors = append(streamInterceptors,
				grpc_retry.StreamClientInterceptor(optsRetry...))
			unaryInterceptors = append(unaryInterceptors,
				grpc_retry.UnaryClientInterceptor(optsRetry...))
		}

		if options.IsEnableMetrics {
			grpc_prometheus.EnableClientHandlingTimeHistogram()
			streamInterceptors = append(streamInterceptors, grpc_prometheus.StreamClientInterceptor)
			unaryInterceptors = append(unaryInterceptors, grpc_prometheus.UnaryClientInterceptor)
		}

		if options.IsEnableSecure {
			opts = append(opts, grpc.WithTransportCredentials(c.Creds))
		} else {
			opts = append(opts, grpc.WithInsecure())
		}

		if options.IsEnableValidator {
			unaryInterceptors = append(unaryInterceptors, grpc_validator.UnaryClientInterceptor())
		}
	}

	sIntOpt := grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
		streamInterceptors...,
	))

	uIntOpt := grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
		unaryInterceptors...,
	))

	opts = append(opts, []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.WaitForReady(false)),
		sIntOpt,
		uIntOpt,
	}...)
	var err error
	opts = append(opts, c.OtherOptions...)
	c.Client, err = grpc.Dial(c.Endpoint, opts...)
	if err != nil {
		return err
	}
	return nil
}

func (c *GrpcClient) Stop(ctx context.Context) error {
	return c.Client.Close()
}
