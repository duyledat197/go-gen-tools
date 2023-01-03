
package services

import (
	"context"
	"time"

	"github.com/duyledat197/interview-hao/internal/models"
	"github.com/duyledat197/interview-hao/internal/repositories"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetListUser(ctx context.Context, offset, limit int) ([]*models.User, error)
	UpdateUser(ctx context.Context, id string, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
