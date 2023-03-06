
package services

import (
	"context"
	"time"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
	
	"github.com/jackc/pgtype"
)

type HubService interface {
	 Create(ctx context.Context, hub *models.Hub) error 
	 GetByID(ctx context.Context, id string) (*models.Hub, error) 
	 GetList(ctx context.Context, offset, limit int) ([]*models.Hub, error) 
	 Update(ctx context.Context, id string, hub *models.Hub) error 
	 Delete(ctx context.Context, id string) error 
}

type hubService struct {
	hubRepo repositories.HubRepository
}

func NewHubService(hubRepo repositories.HubRepository) HubService {
	return &hubService{
		hubRepo: hubRepo,
	}
}

func (s *hubService) Create(ctx context.Context, hub *models.Hub) error {
	if err := s.hubRepo.Create(ctx, hub); err != nil {
		return err
	}
	return nil
}

func (s *hubService) Update(ctx context.Context, id string, hub *models.Hub) error {
	hub.UpdatedAt = pgtype.Timestamptz{
		Time: time.Now(),
	}
	if err := s.hubRepo.Update(ctx, id, hub); err != nil {
		return err
	}
	return nil
}

func (s *hubService) Delete(ctx context.Context, id string) error {
	if err := s.hubRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *hubService) GetList(ctx context.Context, offset, limit int) ([]*models.Hub, error) {
	hubs, err := s.hubRepo.GetList(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return hubs, nil
}

func (s *hubService) GetByID(ctx context.Context, id string) (*models.Hub, error) {
	hub, err := s.hubRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return hub, nil
}
