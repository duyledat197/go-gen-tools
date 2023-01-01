package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	deliveries "github.com/duyledat197/interview-hao/internal/deliveries/grpc"
	"github.com/duyledat197/interview-hao/internal/repositories"
	"github.com/duyledat197/interview-hao/internal/services"
	"github.com/duyledat197/interview-hao/pb"
	"github.com/duyledat197/interview-hao/utils/metadata"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	mgoOption "go.mongodb.org/mongo-driver/mongo/options"
)

type server struct {

	// repo
	userRepo repositories.UserRepository
	teamRepo repositories.TeamRepository

	// service
	userSrv services.UserService
	authSrv services.AuthService
	teamSrv services.TeamService

	// deliveries
	authpb pb.AuthServiceServer
	userpb pb.UserServiceServer
	teampb pb.TeamServiceServer

	// other
	TokenKey string
	mgoDB    *mongo.Database
}

var srv server

func start() error {
	ctx := context.Background()
	if err := srv.loadConfig(); err != nil {
		return err
	}

	if err := srv.connectMongo(); err != nil {
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

func (s *server) connectMongo() error {
	mgoClientOptions := mgoOption.Client().ApplyURI("mongodb+srv://telesale:9bWnTbbchD3g1cnN@cluster0.s979pip.mongodb.net/?retryWrites=true&w=majority")
	// Connect to MongoDB
	var err error
	mgoClient, err := mongo.Connect(context.TODO(), mgoClientOptions)
	if err != nil {
		return err
	}
	s.mgoDB = mgoClient.Database("telesale")
	log.Println("connect mongodb success")
	return nil
}

func (s *server) loadRedis() error {
	// c, err := redis.NewRedisClient(s.cfg.Redis.Address, s.cfg.Redis.Username, s.cfg.Redis.Password)
	// if err != nil {
	// 	return err
	// }
	// s.redisClient = c
	return nil
}

func (s *server) loadRepositories() error {
	// s.userRepo = mongoC.NewUserRepository(s.mgoDB.Collection("user"))
	// s.teamRepo = mongoC.NewTeamRepository(s.mgoDB.Collection("team"))
	return nil
}

func (s *server) loadServices() error {
	s.userSrv = services.NewUserService(s.userRepo)
	s.authSrv = services.NewAuthService(s.userRepo)
	s.teamSrv = services.NewTeamService(s.userRepo, s.teamRepo)
	return nil
}

func (s *server) loadPubSubs() (err error) {
	return nil
}

func (s *server) loadDeliveries() error {
	s.authpb = deliveries.NewAuthDelivery(s.authSrv)
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
	if err := pb.RegisterAuthServiceHandlerServer(ctx, mux, s.authpb); err != nil {
		return err
	}

	if err := pb.RegisterUserServiceHandlerServer(ctx, mux, s.userpb); err != nil {
		return err
	}

	if err := pb.RegisterTeamServiceHandlerServer(ctx, mux, s.teampb); err != nil {
		return err
	}

	port := os.Getenv("PORT")
	log.Printf("server listen on port: %s\n", port)
	handler := cors.AllowAll().Handler(mux)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
