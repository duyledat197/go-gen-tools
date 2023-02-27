package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	deliveries "github.com/duyledat197/go-gen-tools/internal/deliveries/grpc"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
	"github.com/duyledat197/go-gen-tools/internal/repositories/mongo"
	"github.com/duyledat197/go-gen-tools/internal/repositories/postgres"
	"github.com/duyledat197/go-gen-tools/internal/services"
	"github.com/duyledat197/go-gen-tools/pb"
	"github.com/duyledat197/go-gen-tools/utils/logger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type server struct {

	//* repo
	userRepo   repositories.UserRepository
	teamRepo   repositories.TeamRepository
	hubRepo    repositories.HubRepository
	searchRepo repositories.SearchRepository

	//* service
	userSrv   services.UserService
	teamSrv   services.TeamService
	hubSrv    services.HubService
	searchSrv services.SearchService

	//* deliveries
	userpb   pb.UserServiceServer
	teampb   pb.TeamServiceServer
	hubpb    pb.HubServiceServer
	searchpb pb.SearchServiceServer

	//* postgres info
	pgxConfig *pgxpool.Config
	PgxDB     *pgxpool.Pool

	//* config
	config *Config

	//* logger
	logger *zap.Logger
}

var srv server

func start() error {
	ctx := context.Background()
	if err := srv.loadConfig(ctx); err != nil {
		return err
	}

	if err := srv.loadPostgresConnection(ctx); err != nil {
		return err
	}

	if err := srv.loadLogger(); err != nil {
		return err
	}

	if err := srv.loadRepositories(); err != nil {
		return err
	}

	if err := srv.loadServices(); err != nil {
		return err
	}

	if err := srv.loadDeliveries(); err != nil {
		return err
	}

	if err := srv.startServer(ctx); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

func (s *server) loadPostgresConnection(ctx context.Context) error {

	var err error
	s.pgxConfig, err = pgxpool.ParseConfig(s.config.PostgresDB.GetConnectionString())
	if err != nil {
		return err
	}
	s.PgxDB, err = pgxpool.NewWithConfig(ctx, s.pgxConfig)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) loadLogger() error {
	s.logger = logger.NewZapLogger("INFO", true)
	return nil
}
func (s *server) loadRepositories() error {
	// with postgres
	s.userRepo = postgres.NewUserRepository(s.PgxDB)
	s.teamRepo = postgres.NewTeamRepository(s.PgxDB)
	s.hubRepo = postgres.NewHubRepository(s.PgxDB)

	// with mongo
	s.searchRepo = mongo.NewSearchRepository(s.PgxDB)

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

func (s *server) startServer(ctx context.Context) error {

	httpPort := s.config.HTTP.Port
	grpcPort := s.config.GRPC.Port

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		//* middleware
		mux := runtime.NewServeMux(
		// runtime.WithMetadata(metadata.Authentication),
		)

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

		handler := cors.AllowAll().Handler(mux)
		s.logger.Sugar().Infoln("HTTP Server listens on port: %s\n", httpPort)
		if err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), handler); err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {

		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		if err != nil {
			return err
		}

		grpcServer := grpc.NewServer()
		pb.RegisterUserServiceServer(grpcServer, s.userpb)
		pb.RegisterTeamServiceServer(grpcServer, s.teampb)
		pb.RegisterHubServiceServer(grpcServer, s.hubpb)
		pb.RegisterSearchServiceServer(grpcServer, s.searchpb)

		s.logger.Sugar().Infoln("GRPC Server listens on port: %v", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		s.logger.Sugar().Fatalln("start server failed:", err)
	}

	return nil
}
