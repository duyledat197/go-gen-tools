
package services

import (
	"context"
	"time"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
	
	"github.com/jackc/pgtype"
)

type UserService interface {
	 Create(ctx context.Context, user *models.User) error 
	 GetByID(ctx context.Context, id string) (*models.User, error) 
	 GetList(ctx context.Context, offset, limit int) ([]*models.User, error) 
	 Update(ctx context.Context, id string, user *models.User) error 
	 Delete(ctx context.Context, id string) error 
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Create(ctx context.Context, user *models.User) error {
	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *userService) Update(ctx context.Context, id string, user *models.User) error {
	user.UpdatedAt = pgtype.Timestamptz{
		Time: time.Now(),
	}
	if err := s.userRepo.Update(ctx, id, user); err != nil {
		return err
	}
	return nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	if err := s.userRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *userService) GetList(ctx context.Context, offset, limit int) ([]*models.User, error) {
	users, err := s.userRepo.GetList(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetByID(ctx context.Context, id string) (*models.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
