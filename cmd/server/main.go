package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	grpcServer *grpc.Server
	httpServer *http.Server
}

var srv server

func start(ctx context.Context) error {

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

func stop(ctx context.Context) error {
	if err := srv.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	srv.grpcServer.GracefulStop()
	return nil
}

func main() {
	ctx := context.Background()
	// ?
	timeWait := 15 * time.Second
	signChan := make(chan os.Signal, 1)

	if err := start(ctx); err != nil {
		log.Fatal(err)
	}
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
	<-signChan
	log.Println("Shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), timeWait)
	defer func() {
		log.Println("Close another connection")
		cancel()
	}()
	if err := stop(ctx); err == context.DeadlineExceeded {
		log.Print("Halted active connections")
	}
	close(signChan)
	log.Printf("Server down Completed")
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
		s.httpServer = &http.Server{
			Addr:    fmt.Sprintf(":%s", httpPort),
			Handler: handler,
		}
		s.logger.Sugar().Infof("HTTP Server listens on port: %s\n", httpPort)
		if err := s.httpServer.ListenAndServe(); err != nil {
			return err
		}

		return nil
	})

	eg.Go(func() error {

		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		if err != nil {
			return err
		}

		s.grpcServer = grpc.NewServer()
		pb.RegisterUserServiceServer(s.grpcServer, s.userpb)
		pb.RegisterTeamServiceServer(s.grpcServer, s.teampb)
		pb.RegisterHubServiceServer(s.grpcServer, s.hubpb)
		pb.RegisterSearchServiceServer(s.grpcServer, s.searchpb)

		s.logger.Sugar().Infoln("GRPC Server listens on port: %v", grpcPort)
		if err := s.grpcServer.Serve(lis); err != nil {
			return err
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		s.logger.Sugar().Fatalln("start server failed:", err)
	}

	return nil
}
