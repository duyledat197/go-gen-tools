package deliveries

import (
	"context"
	"fmt"

	"github.com/duyledat197/go-gen-tools/idl/pb"
	"github.com/duyledat197/go-gen-tools/internal/services"
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
	if err := d.teamService.Create(ctx, transform.PbToTeamPtr(req.GetTeam())); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Errorf("Create: %v", err).Error())
	}
	return &pb.CreateTeamResponse{}, nil
}

func (d *teamDelivery) UpdateTeam(ctx context.Context, req *pb.UpdateTeamRequest) (*pb.UpdateTeamResponse, error) {
	if err := d.teamService.Update(ctx, req.Team.Id, transform.PbToTeamPtr(req.GetTeam())); err != nil {
		return nil, err
	}
	return &pb.UpdateTeamResponse{
		Success: true,
	}, nil
}

func (d *teamDelivery) DeleteTeam(ctx context.Context, req *pb.DeleteTeamRequest) (*pb.DeleteTeamResponse, error) {
	if err := d.teamService.Delete(ctx, req.Team.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteTeamResponse{
		Success: true,
	}, nil
}

func (d *teamDelivery) GetList(ctx context.Context, req *pb.GetListTeamRequest) (*pb.GetListTeamResponse, error) {
	teams, err := d.teamService.GetList(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	return &pb.GetListTeamResponse{
		Data: transform.TeamToPbPtrList(teams),
	}, nil
}

func (d *teamDelivery) GetTeamByID(ctx context.Context, req *pb.GetTeamByIDRequest) (*pb.GetTeamByIDResponse, error) {
	team, err := d.teamService.GetByID(ctx, req.GetTeamID())
	if err != nil {
		return nil, err
	}
	return &pb.GetTeamByIDResponse{
		Data: transform.TeamToPbPtr(team),
	}, nil
}
