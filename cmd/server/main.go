package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	deliveries "github.com/duyledat197/go-gen-tools/internal/deliveries/grpc"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
	"github.com/duyledat197/go-gen-tools/internal/services"
	"github.com/duyledat197/go-gen-tools/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

type server struct {

	// repo
	userRepo   repositories.UserRepository
	teamRepo   repositories.TeamRepository
	hubRepo    repositories.HubRepository
	searchRepo repositories.SearchRepository

	// service
	userSrv   services.UserService
	teamSrv   services.TeamService
	hubSrv    services.HubService
	searchSrv services.SearchService

	// deliveries
	userpb   pb.UserServiceServer
	teampb   pb.TeamServiceServer
	hubpb    pb.HubServiceServer
	searchpb pb.SearchServiceServer

	// other
	db *sql.DB
}

var srv server

func start() error {
	ctx := context.Background()
	if err := srv.loadConfig(); err != nil {
		return err
	}

	if err := srv.loadDB(); err != nil {
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

	srv.startServer(ctx)
	return nil
}

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

func (s *server) loadDB() error {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *server) loadRepositories() error {
	s.userRepo = repositories.NewUserRepository(s.db)
	s.teamRepo = repositories.NewTeamRepository(s.db)
	s.hubRepo = repositories.NewHubRepository(s.db)
	s.searchRepo = repositories.NewSearchRepository(s.db)

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

func (s *server) loadConfig() error {
	return nil
}

func (s *server) startServer(ctx context.Context) error {
	var serverError = make(chan error)
	var waitGroup sync.WaitGroup

	httpPort := "8585"
	grpcPort := "5000"

	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()

		mux := runtime.NewServeMux(
		// runtime.WithMetadata(metadata.Authentication),
		)

		if err := pb.RegisterUserServiceHandlerServer(ctx, mux, s.userpb); err != nil {
			serverError <- err
		}
		if err := pb.RegisterTeamServiceHandlerServer(ctx, mux, s.teampb); err != nil {
			serverError <- err
		}
		if err := pb.RegisterHubServiceHandlerServer(ctx, mux, s.hubpb); err != nil {
			serverError <- err
		}
		if err := pb.RegisterSearchServiceHandlerServer(ctx, mux, s.searchpb); err != nil {
			serverError <- err
		}

		handler := cors.AllowAll().Handler(mux)
		log.Printf("HTTP Server listens on port: %s\n", httpPort)
		http.ListenAndServe(fmt.Sprintf(":%s", httpPort), handler)
	}()

	go func() {
		defer waitGroup.Done()

		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		if err != nil {
			serverError <- err
		}

		grpcServer := grpc.NewServer()
		pb.RegisterUserServiceServer(grpcServer, s.userpb)
		pb.RegisterTeamServiceServer(grpcServer, s.teampb)
		pb.RegisterHubServiceServer(grpcServer, s.hubpb)
		pb.RegisterSearchServiceServer(grpcServer, s.searchpb)

		log.Printf("GRPC Server listens on port: %v", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			serverError <- err
		}
	}()

	for <-serverError != nil {
		log.Fatal("start server failed:", <-serverError)
	}

	waitGroup.Wait()

	return nil
}
