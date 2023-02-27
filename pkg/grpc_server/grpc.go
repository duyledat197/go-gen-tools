package grpc_server

import (
	"context"
	"expvar"
	"fmt"
	"net"
	"net/http"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/duyledat197/go-gen-tools/pkg/ratelimit"
	"github.com/duyledat197/go-gen-tools/pkg/registry"
	"github.com/duyledat197/go-gen-tools/pkg/tracing"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Options struct {
	//* for enable metrics with prometheus
	IsEnableMetrics bool

	//* for enable client load balancer with consul
	IsEnableClientLoadBalancer bool

	//* for enable tracing
	IsEnableTracer bool

	//* for enable ctx tags
	IsEnableCtxTags bool

	//* for enable zap logger
	IsEnableLogger bool

	//* for enable recovery
	IsEnableRecovery bool

	//* for enable auth function (authenticate layer)
	IsEnableAuthFunc bool

	//* for start prometheus server
	IsEnablePrometheusServer bool

	//* for enable ratelimit
	IsEnableRateLimit bool

	//* for enable validator
	IsEnableValidator bool
}

type GrpcServer struct {
	ServiceName    string
	Consul         *registry.ConsulRegister
	Tracer         *tracing.TracerClient
	AuthFunction   grpc_auth.AuthFunc
	server         *grpc.Server
	Logger         *zap.Logger
	Address        *config.ConnectionAddr
	Options        *Options
	MaxMessageSize int //* default = 0 mean 4MB
	Handlers       func(ctx context.Context, server *grpc.Server) error

	OtherOptions []grpc.ServerOption
}

func (s *GrpcServer) Init(ctx context.Context) error {
	filterFn := grpc_opentracing.WithFilterFunc(func(ctx context.Context, fullMethodName string) bool {
		return fullMethodName != "/grpc.health.v1.Health/Check"
	})

	recoveryFn := func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	var (
		streamInterceptors []grpc.StreamServerInterceptor
		unaryInterceptors  []grpc.UnaryServerInterceptor
		opts               []grpc.ServerOption
	)

	if s.Options != nil {
		options := s.Options
		if options.IsEnableRecovery {
			streamInterceptors = append(streamInterceptors, grpc_recovery.StreamServerInterceptor(
				grpc_recovery.WithRecoveryHandler(recoveryFn)))
			unaryInterceptors = append(unaryInterceptors, grpc_recovery.UnaryServerInterceptor(
				grpc_recovery.WithRecoveryHandler(recoveryFn)))
		}

		if options.IsEnableCtxTags {
			streamInterceptors = append(streamInterceptors, grpc_ctxtags.StreamServerInterceptor())
			unaryInterceptors = append(unaryInterceptors, grpc_ctxtags.UnaryServerInterceptor())
		}

		if options.IsEnableTracer {
			streamInterceptors = append(streamInterceptors,
				grpc_opentracing.StreamServerInterceptor(
					grpc_opentracing.WithTracer(s.Tracer.Tracer),
					filterFn,
				),
			)
			unaryInterceptors = append(unaryInterceptors,
				grpc_opentracing.UnaryServerInterceptor(
					grpc_opentracing.WithTracer(s.Tracer.Tracer),
					filterFn),
			)
		}

		if options.IsEnableMetrics {
			streamInterceptors = append(streamInterceptors, grpc_prometheus.StreamServerInterceptor)
			unaryInterceptors = append(unaryInterceptors, grpc_prometheus.UnaryServerInterceptor)
		}

		if options.IsEnableAuthFunc {
			streamInterceptors = append(streamInterceptors, grpc_auth.StreamServerInterceptor(s.AuthFunction))
			unaryInterceptors = append(unaryInterceptors, grpc_auth.UnaryServerInterceptor(s.AuthFunction))
		}

		if options.IsEnableClientLoadBalancer {
			if err := s.Consul.Register(); err != nil {
				s.Logger.Panic("connect consul server error:", zap.Error(err))
			}
		}

		if options.IsEnableLogger {
			streamInterceptors = append(streamInterceptors, grpc_zap.StreamServerInterceptor(s.Logger))
			unaryInterceptors = append(unaryInterceptors, grpc_zap.UnaryServerInterceptor(s.Logger))
		}

		if options.IsEnablePrometheusServer {
			grpc_prometheus.Register(s.server)
			mux := http.NewServeMux()

			mux.Handle("/debug/vars", expvar.Handler())
			mux.Handle("/metrics", promhttp.Handler())

			go func() {
				if err := http.ListenAndServe(fmt.Sprintf(":9090"), mux); err != nil {
					s.Logger.Panic("start prometheus server error:", zap.Error(err))
				}
			}()
		}

		if options.IsEnableRateLimit {
			limiter := ratelimit.NewLimitter(3, 10)
			streamInterceptors = append(streamInterceptors, grpc_ratelimit.StreamServerInterceptor(limiter))
			unaryInterceptors = append(unaryInterceptors, grpc_ratelimit.UnaryServerInterceptor(limiter))
		}

		if options.IsEnableValidator {
			streamInterceptors = append(streamInterceptors, grpc_validator.StreamServerInterceptor())
			unaryInterceptors = append(unaryInterceptors, grpc_validator.UnaryServerInterceptor())
		}
	}

	if s.MaxMessageSize != 0 {
		opts = append(opts,
			grpc.MaxRecvMsgSize(s.MaxMessageSize),
			grpc.MaxSendMsgSize(s.MaxMessageSize),
			grpc.MaxMsgSize(s.MaxMessageSize),
		)
	}
	opts = append(opts,
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			streamInterceptors...,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryInterceptors...,
		)))
	opts = append(opts, s.OtherOptions...)
	s.server = grpc.NewServer(
		opts...,
	)
	s.Handlers(ctx, s.server)
	return nil
}

func (s *GrpcServer) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Address.Port))
	if err != nil {
		return err
	}
	s.Logger.Sugar().Infoln("GRPC Server listens on port: %v", s.Address.Port)
	if err := s.server.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *GrpcServer) Stop(ctx context.Context) error {
	s.server.GracefulStop()
	s.Consul.Deregister()
	return nil
}
