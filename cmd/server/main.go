package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	deliveries "github.com/duyledat197/interview-hao/internal/deliveries/grpc"
	"github.com/duyledat197/interview-hao/internal/models"
	"github.com/duyledat197/interview-hao/internal/repositories"
	"github.com/duyledat197/interview-hao/internal/services"
	"github.com/duyledat197/interview-hao/pb"
	"github.com/duyledat197/interview-hao/utils/metadata"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {

	// repo
	userRepo repositories.UserRepository

	// service
	userSrv services.UserService

	// deliveries
	userpb pb.UserServiceServer

	// other
	TokenKey string
	mgoDB    *mongo.Database
	db       *sql.DB
	queries  *models.Queries
}

var srv server

func start() error {
	ctx := context.Background()
	if err := srv.loadConfig(); err != nil {
		return err
	}

	if err := srv.connectPsql(); err != nil {
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

	srv.startGRPCServer(ctx)
	return nil
}

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

func (s *server) NewQueries() error {
	s.queries = models.New(s.db)
	return nil
}

func (s *server) connectPsql() error {
	db, err := sql.Open("pgx", os.Getenv("postgres://postgres:password@localhost/cubicasa?sslmode=disable"))
	if err != nil {
		return err
	}
	s.db = db
	defer db.Close()

	return nil
}

func (s *server) loadRepositories() error {
	s.userRepo = repositories.NewUserRepository(s.queries)
	return nil
}

func (s *server) loadServices() error {
	s.userSrv = services.NewUserService(s.userRepo)
	return nil
}

func (s *server) loadPubSubs() (err error) {
	return nil
}

func (s *server) loadDeliveries() error {
	s.userpb = deliveries.NewUserDelivery(s.userSrv)
	return nil
}

func (s *server) loadLogger() error {
	return nil
}

func (s *server) loadConfig() error {
	return nil
}

func (s *server) startGRPCServer(ctx context.Context) error {
	mux := runtime.NewServeMux(
		runtime.WithMetadata(metadata.Authentication),
	)

	if err := pb.RegisterUserServiceHandlerServer(ctx, mux, s.userpb); err != nil {
		return err
	}

	port := os.Getenv("PORT")
	log.Printf("server listen on port: %s\n", port)
	handler := cors.AllowAll().Handler(mux)

	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
