package deliveries

import (
	"context"
	"fmt"

	"github.com/duyledat197/go-gen-tools/internal/services"
	"github.com/duyledat197/go-gen-tools/pb"
	"github.com/duyledat197/go-gen-tools/transform"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type teamDelivery struct {
	teamService services.TeamService
	pb.UnimplementedTeamServiceServer
}

func NewTeamDelivery(teamService services.TeamService) pb.TeamServiceServer {
	return &teamDelivery{
		teamService: teamService,
	}
}

func (d *teamDelivery) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.CreateTeamResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("validate failed: %w", err).Error())
	}
	if err := d.teamService.Create(ctx, transform.PbToTeamPtr(req.GetTeam())); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Errorf("Create: %v", err).Error())
	}
	return &pb.CreateTeamResponse{}, nil
}
