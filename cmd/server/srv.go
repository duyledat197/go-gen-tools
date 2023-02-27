package main

import (
	"context"

	"github.com/duyledat197/go-gen-tools/config"
	deliveries "github.com/duyledat197/go-gen-tools/internal/deliveries/grpc"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
	"github.com/duyledat197/go-gen-tools/internal/repositories/mongo"
	"github.com/duyledat197/go-gen-tools/internal/repositories/postgres"
	"github.com/duyledat197/go-gen-tools/internal/services"
	"github.com/duyledat197/go-gen-tools/pb"
	"github.com/duyledat197/go-gen-tools/pkg/grpc_client"
	"github.com/duyledat197/go-gen-tools/pkg/grpc_server"
	"github.com/duyledat197/go-gen-tools/pkg/http_server"
	"github.com/duyledat197/go-gen-tools/pkg/kafka_utils"
	"github.com/duyledat197/go-gen-tools/pkg/mongo_client"
	"github.com/duyledat197/go-gen-tools/pkg/postgres_client"
	"github.com/duyledat197/go-gen-tools/pkg/registry"
	"github.com/duyledat197/go-gen-tools/pkg/tracing"
	"github.com/duyledat197/go-gen-tools/utils/authenticate"
	"github.com/duyledat197/go-gen-tools/utils/logger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	//* authenticator
	authenticator authenticate.Authenticator

	//* reposities
	userRepo   repositories.UserRepository
	teamRepo   repositories.TeamRepository
	hubRepo    repositories.HubRepository
	searchRepo repositories.SearchRepository

	//* services
	userSrv   services.UserService
	teamSrv   services.TeamService
	hubSrv    services.HubService
	searchSrv services.SearchService

	//* deliveries
	userpb   pb.UserServiceServer
	teampb   pb.TeamServiceServer
	hubpb    pb.HubServiceServer
	searchpb pb.SearchServiceServer

	//* grpc clients
	userClient *grpc_client.GrpcClient
	teamClient *grpc_client.GrpcClient
	hubClient  *grpc_client.GrpcClient

	//* postgres info
	postgresClient *postgres_client.PostgresClient

	//* mongo info
	mongoClient *mongo_client.MongoClient

	//* config
	config *config.Config

	//* logger
	logger *zap.Logger

	//* servers
	grpcServer *grpc_server.GrpcServer
	httpServer *http_server.HttpServer

	//* third_party services
	consul     *registry.ConsulClient
	tracer     *tracing.TracerClient
	publisher  *kafka_utils.Publisher
	subscriber *kafka_utils.Subscriber

	processors []processor
	databases  []database
}

type processor interface {
	Init(ctx context.Context) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type database interface {
	Connect(ctx context.Context) error
	Stop(ctx context.Context) error
}

func (s *server) loadDatabaseClients(ctx context.Context) error {
	if err := srv.loadPostgresClient(ctx); err != nil {
		return err
	}

	if err := srv.loadMongoClient(ctx); err != nil {
		return err
	}

	s.databases = append(s.databases, s.postgresClient, s.mongoClient)
	return nil
}

func (s *server) loadPostgresClient(ctx context.Context) error {
	s.postgresClient = &postgres_client.PostgresClient{
		Database: s.config.PostgresDB,
		Logger:   s.logger,
		Options:  &postgres_client.Options{},
	}
	return nil
}

func (s *server) loadMongoClient(ctx context.Context) error {
	s.mongoClient = &mongo_client.MongoClient{
		Database: s.config.MongoDB,
		Logger:   s.logger,
		Options:  &mongo_client.Options{},
	}
	return nil
}

func (s *server) loadLogger() error {
	s.logger = logger.NewZapLogger("INFO", true)
	return nil
}
func (s *server) loadRepositories() error {
	// with postgres
	s.userRepo = postgres.NewUserRepository(s.postgresClient.Pool)
	s.teamRepo = postgres.NewTeamRepository(s.postgresClient.Pool)
	s.hubRepo = postgres.NewHubRepository(s.postgresClient.Pool)

	// with mongo
	s.searchRepo = mongo.NewSearchRepository(s.postgresClient.Pool)

	return nil
}

func (s *server) loadServices() error {
	s.userSrv = services.NewUserService(s.userRepo)
	s.teamSrv = services.NewTeamService(s.teamRepo)
	s.hubSrv = services.NewHubService(s.hubRepo)
	s.searchSrv = services.NewSearchService(s.searchRepo)

	return nil
}

func (s *server) loadDeliveries() error {
	s.userpb = deliveries.NewUserDelivery(s.userSrv)
	s.teampb = deliveries.NewTeamDelivery(s.teamSrv)
	s.hubpb = deliveries.NewHubDelivery(s.hubSrv)
	s.searchpb = deliveries.NewSearchDelivery(s.searchSrv)

	return nil
}

func (s *server) loadConfig(ctx context.Context) error {
	return nil
}

func (s *server) loadClients(ctx context.Context) error {
	defaultOptions := &grpc_client.Options{
		IsEnableHystrix:            true,
		IsEnableClientLoadBalancer: true,
		IsEnableTracing:            true,
		IsEnableRetry:              true,
		IsEnableMetrics:            true,
		IsEnableSecure:             false,
		IsEnableValidator:          true,
	}
	//* load grpc clients
	s.teamClient = &grpc_client.GrpcClient{
		ServiceName: "Team",
		Consul:      s.consul,
		Tracer:      s.tracer,
		Options:     defaultOptions,
	}

	s.userClient = &grpc_client.GrpcClient{
		ServiceName: "User",
		Consul:      s.consul,
		Tracer:      s.tracer,
		Options:     defaultOptions,
	}

	s.hubClient = &grpc_client.GrpcClient{
		ServiceName: "Hub",
		Consul:      s.consul,
		Tracer:      s.tracer,
		Options:     defaultOptions,
	}

	//* load http clients

	s.databases = append(s.databases, s.teamClient, s.userClient, s.hubClient)

	return nil
}

func (s *server) loadServers(ctx context.Context) error {
	s.httpServer = &http_server.HttpServer{
		Address:       s.config.HTTP,
		Logger:        s.logger,
		Authenticator: s.authenticator,
		Handlers: func(ctx context.Context, mux *runtime.ServeMux) error {
			if err := pb.RegisterUserServiceHandlerServer(ctx, mux, s.userpb); err != nil {
				return err
			}
			if err := pb.RegisterTeamServiceHandlerServer(ctx, mux, s.teampb); err != nil {
				return err
			}
			if err := pb.RegisterHubServiceHandlerServer(ctx, mux, s.hubpb); err != nil {
				return err
			}
			if err := pb.RegisterSearchServiceHandlerServer(ctx, mux, s.searchpb); err != nil {
				return err
			}
			return nil
		},
		Options: &http_server.Options{},
	}

	s.grpcServer = &grpc_server.GrpcServer{
		ServiceName: s.config.ServiceName,
		Tracer:      s.tracer,
		Address:     s.config.GRPC,
		Logger:      s.logger,
		Handlers: func(ctx context.Context, server *grpc.Server) error {
			pb.RegisterUserServiceServer(server, s.userpb)
			pb.RegisterTeamServiceServer(server, s.teampb)
			pb.RegisterHubServiceServer(server, s.hubpb)
			pb.RegisterSearchServiceServer(server, s.searchpb)
			return nil
		},
		Options: &grpc_server.Options{
			IsEnableMetrics:            true,
			IsEnableTracer:             true,
			IsEnableClientLoadBalancer: true,
			IsEnableCtxTags:            true,
			IsEnableLogger:             true,
			IsEnableRecovery:           true,
			IsEnableAuthFunc:           true,
			IsEnablePrometheusServer:   true,
			IsEnableRateLimit:          true,
			IsEnableValidator:          true,
		},
	}

	s.processors = append(s.processors, s.grpcServer, s.httpServer)

	return nil
}

func (s *server) loadThirdPartyClients(ctx context.Context) error {
	s.consul = &registry.ConsulClient{
		Address: s.config.Consul,
		Logger:  s.logger,
		Options: &registry.Options{},
	}

	s.tracer = &tracing.TracerClient{
		ServiceName: s.config.ServiceName,
		Address:     s.config.Tracer,
		Logger:      s.logger,
	}

	s.databases = append(s.databases, s.consul, s.tracer)
	return nil
}
