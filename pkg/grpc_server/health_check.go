package grpc_server

import (
	"context"

	pb "google.golang.org/grpc/health/grpc_health_v1"
)

// NewHealthService ...
func NewHealthService() pb.HealthServer {
	return &healthService{}
}

// healthService ...
type healthService struct{}

func (s *healthService) Check(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

func (s *healthService) Watch(*pb.HealthCheckRequest, pb.Health_WatchServer) error {
	return nil
}
