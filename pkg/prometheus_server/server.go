package prometheus_server

import (
	"context"
	"expvar"
	"fmt"
	"net/http"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type PrometheusServer struct {
	server  *http.Server
	Logger  *zap.Logger
	Address *config.ConnectionAddr
}

func (s *PrometheusServer) Init(ctx context.Context) error {
	mux := http.NewServeMux()

	mux.Handle("/debug/vars", expvar.Handler())
	mux.Handle("/metrics", promhttp.Handler())
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Address.Port),
		Handler: mux,
	}

	return nil
}

func (s *PrometheusServer) Start(ctx context.Context) error {
	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("start prometheus server error: %w", err)
	}
	return nil
}

func (s *PrometheusServer) Stop(ctx context.Context) error {
	return s.server.Close()
}
