package deliveries

import (
	"context"
	"fmt"

	"github.com/duyledat197/interview-hao/internal/services"
	"github.com/duyledat197/interview-hao/pb"
	"github.com/duyledat197/interview-hao/transform"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	if err := d.teamService.CreateTeam(ctx, transform.PbToTeamPtr(req.GetTeam())); err != nil {
		return nil, err
	}
	return &pb.CreateTeamResponse{}, nil
}

func (d *teamDelivery) AddUsersToTeam(ctx context.Context, req *pb.AddUsersToTeamRequest) (*pb.AddUsersToTeamResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.TeamId)
	if err != nil {
		return nil, err
	}
	var userIDs []primitive.ObjectID
	for _, userID := range req.UserIds {
		id, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, id)
	}
	if err := d.teamService.AddUsersToTeam(ctx, id, userIDs); err != nil {
		return nil, err
	}
	return &pb.AddUsersToTeamResponse{
		Success: true,
	}, nil
}

func (d *teamDelivery) GetTeams(ctx context.Context, req *pb.GetTeamsRequest) (*pb.GetTeamsResponse, error) {
	team, total, err := d.teamService.GetListTeam(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	result := transform.TeamToPbPtrList(team)
	return &pb.GetTeamsResponse{
		Data:  result,
		Total: total,
	}, nil
}

func (d *teamDelivery) GetTeamByID(ctx context.Context, req *pb.GetTeamByIDRequest) (*pb.GetTeamByIDResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.TeamId)
	if err != nil {
		return nil, err
	}
	team, err := d.teamService.GetTeamByID(ctx, id)
	return &pb.GetTeamByIDResponse{
		Team: transform.TeamToPbPtr(team),
	}, nil
}

func (d *teamDelivery) UpdateTeam(ctx context.Context, req *pb.UpdateTeamRequest) (*pb.UpdateTeamResponse, error) {
	if err := d.teamService.Update(ctx, transform.PbToTeamPtr(req.Team)); err != nil {
		return nil, err
	}
	return &pb.UpdateTeamResponse{
		Success: true,
	}, nil
}
