package services

import (
	"context"
	"time"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
	"github.com/jackc/pgtype"
)

type TeamService interface {
	Create(ctx context.Context, team *models.Team) error
	GetByID(ctx context.Context, id string) (*models.Team, error)
	GetList(ctx context.Context, offset, limit int) ([]*models.Team, error)
	Update(ctx context.Context, id string, team *models.Team) error
	Delete(ctx context.Context, id string) error
}

type teamService struct {
	teamRepo repositories.TeamRepository
}

func NewTeamService(teamRepo repositories.TeamRepository) TeamService {
	return &teamService{
		teamRepo: teamRepo,
	}
}

func (s *teamService) Create(ctx context.Context, team *models.Team) error {
	if err := s.teamRepo.Create(ctx, team); err != nil {
		return err
	}
	return nil
}

func (s *teamService) Update(ctx context.Context, id string, team *models.Team) error {
	team.UpdatedAt = pgtype.Timestamptz{
		Time: time.Now(),
	}
	if err := s.teamRepo.Update(ctx, id, team); err != nil {
		return err
	}
	return nil
}

func (s *teamService) Delete(ctx context.Context, id string) error {
	if err := s.teamRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *teamService) GetList(ctx context.Context, offset, limit int) ([]*models.Team, error) {
	teams, err := s.teamRepo.GetList(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (s *teamService) GetByID(ctx context.Context, id string) (*models.Team, error) {
	team, err := s.teamRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return team, nil
}
