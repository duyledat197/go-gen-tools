package services

import (
	"context"
	"time"

	"github.com/duyledat197/interview-hao/internal/models"
	"github.com/duyledat197/interview-hao/internal/repositories"
	"github.com/duyledat197/interview-hao/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamService interface {
	CreateTeam(ctx context.Context, team *models.Team) error
	GetListTeam(ctx context.Context, offset, limit int) ([]*models.Team, int64, error)
	AddUsersToTeam(ctx context.Context, teamID primitive.ObjectID, userIDs []primitive.ObjectID) error
	GetTeamByID(ctx context.Context, teamID primitive.ObjectID) (*models.Team, error)
	Update(ctx context.Context, team *models.Team) error
}

type teamService struct {
	userRepo repositories.UserRepository
	teamRepo repositories.TeamRepository
}

func (s *teamService) CreateTeam(ctx context.Context, team *models.Team) error {
	adminID, err := Authorized(ctx, s.userRepo, []string{
		pb.UserRole_SUPER_ADMIN.String(),
		pb.UserRole_ADMIN.String(),
		pb.UserRole_MANAGER.String(),
	})
	if err != nil {
		return err
	}
	now := time.Now()
	team.CreatedBy = adminID
	team.CreatedAt = now
	team.UpdatedAt = now
	if err := s.teamRepo.Create(ctx, team); err != nil {
		return err
	}
	return nil
}

func (s *teamService) GetListTeam(ctx context.Context, offset int, limit int) ([]*models.Team, int64, error) {
	userID, role, err := Authorized2(ctx, s.userRepo, []string{
		pb.UserRole_SUPER_ADMIN.String(),
		pb.UserRole_ADMIN.String(),
		pb.UserRole_MANAGER.String(),
	})
	if err != nil {
		return nil, 0, err
	}
	filter := &models.Team{}
	if role == pb.UserRole_MANAGER.String() {
		filter.CreatedBy = userID
	}
	teams, err := s.teamRepo.FindByFilter(ctx, filter, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.teamRepo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return teams, total, nil
}

func (s *teamService) AddUsersToTeam(ctx context.Context, teamID primitive.ObjectID, userIDs []primitive.ObjectID) error {
	userID, role, err := Authorized2(ctx, s.userRepo, []string{
		pb.UserRole_SUPER_ADMIN.String(),
		pb.UserRole_ADMIN.String(),
		pb.UserRole_MANAGER.String(),
	})
	if err != nil {
		return err
	}
	filter := &models.Team{
		ID: teamID,
	}
	if role == pb.UserRole_MANAGER.String() {
		filter.CreatedBy = userID
	}
	if err := s.teamRepo.UpdateUsers(ctx, filter, userIDs); err != nil {
		return err
	}
	return nil
}

func (s *teamService) GetTeamByID(ctx context.Context, teamID primitive.ObjectID) (*models.Team, error) {
	result, err := s.teamRepo.FindByID(ctx, teamID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *teamService) Update(ctx context.Context, team *models.Team) error {
	if _, err := Authorized(ctx, s.userRepo, []string{pb.UserRole_SUPER_ADMIN.String(), pb.UserRole_ADMIN.String()}); err != nil {
		return err
	}

	if err := s.teamRepo.Update(ctx, team); err != nil {
		return err
	}
	return nil
}
func NewTeamService(userRepo repositories.UserRepository, teamRepo repositories.TeamRepository) TeamService {
	return &teamService{
		userRepo: userRepo,
		teamRepo: teamRepo,
	}
}
