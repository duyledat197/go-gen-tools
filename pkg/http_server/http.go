package http_server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/duyledat197/go-gen-tools/utils/authenticate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
)

type Options struct{}
type HttpServer struct {
	DocFileName   string
	DocFileRoot   string
	Address       *config.ConnectionAddr
	server        *http.Server
	Handlers      func(ctx context.Context, mux *runtime.ServeMux) error
	Logger        *zap.Logger
	Authenticator authenticate.Authenticator
	Options       *Options
}

func (s *HttpServer) Init(ctx context.Context) error {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions:   protojson.MarshalOptions{UseEnumNumbers: false},
			UnmarshalOptions: protojson.UnmarshalOptions{AllowPartial: true},
		}),
		runtime.WithMetadata(MapMetaDataWithBearerToken(s.Authenticator)),
	)
	if err := s.Handlers(ctx, mux); err != nil {
		return err
	}

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Address.Port),
		Handler: AllowCORS(mux),
	}
	return nil
}

func (s *HttpServer) Start(ctx context.Context) error {
	s.Logger.Sugar().Infof("HTTP Server listens on port: %s\n", s.Address.Port)
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *HttpServer) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
