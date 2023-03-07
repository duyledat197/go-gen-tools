package swagger_server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/duyledat197/go-gen-tools/utils/pathutils"

	swaggerui "github.com/esurdam/go-swagger-ui"
	"go.uber.org/zap"
)

type SwaggerServer struct {
	Address *config.ConnectionAddr
	server  *http.Server
	Logger  *zap.Logger
}

func (s *SwaggerServer) Init(ctx context.Context) error {
	root := pathutils.GetPkgDir()
	mux := swaggerui.NewServeMuxWithRoot(func(str string) ([]byte, error) {
		file, err := ioutil.ReadFile(filepath.Join(fmt.Sprintf("../../%s/%s", root, str)))
		if err != nil {
			return nil, err
		}
		return file, nil
	}, "docs/swagger", root)
	s.server = &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%s", s.Address.Port),
	}
	return nil
}

func (s *SwaggerServer) Start(ctx context.Context) error {
	s.Logger.Sugar().Infof("swagger server start in port: %w\n", s.Address.Port)
	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("start swagger server error: %w", err)
	}
	return nil
}

func (s *SwaggerServer) Stop(ctx context.Context) error {
	return s.server.Close()
}
