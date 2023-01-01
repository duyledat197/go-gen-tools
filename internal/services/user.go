package services

import (
	"context"
	"fmt"
	"time"

	"github.com/lalaland/backend/internal/models"
	"github.com/lalaland/backend/internal/repositories"
	"github.com/lalaland/backend/pb"
	"github.com/lalaland/backend/utils/metadata"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/slices"
)

// Service interface
type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, userID primitive.ObjectID) (*models.User, error)
	GetListUser(ctx context.Context, offset, limit int) ([]*models.User, int64, error)
	UpdateUser(ctx context.Context, userID primitive.ObjectID, user *models.User) error
	DeleteUser(ctx context.Context, userID primitive.ObjectID) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func Authorized(ctx context.Context, userRepo repositories.UserRepository, acceptRoles []string) (primitive.ObjectID, error) {
	adminID, ok := metadata.GetUserID(ctx)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("need admin token")
	}
	adID, err := primitive.ObjectIDFromHex(adminID)
	if err != nil {
		return primitive.NilObjectID, err
	}
	admin, err := userRepo.FindByUserID(ctx, adID)
	if err != nil {
		return primitive.NilObjectID, err
	}
	if !slices.Contains(acceptRoles, admin.Role) {
		return primitive.NilObjectID, fmt.Errorf("doesn't accept role")
	}
	return admin.ID, nil
}

func Authorized2(ctx context.Context, userRepo repositories.UserRepository, acceptRoles []string) (primitive.ObjectID, string, error) {
	adminID, ok := metadata.GetUserID(ctx)
	if !ok {
		return primitive.NilObjectID, "", fmt.Errorf("need admin token")
	}
	adID, err := primitive.ObjectIDFromHex(adminID)
	if err != nil {
		return primitive.NilObjectID, "", err
	}
	admin, err := userRepo.FindByUserID(ctx, adID)
	if err != nil {
		return primitive.NilObjectID, "", err
	}
	if !slices.Contains(acceptRoles, admin.Role) {
		return primitive.NilObjectID, "", fmt.Errorf("doesn't accept role")
	}
	return admin.ID, admin.Role, nil
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	adminID, err := Authorized(ctx, s.userRepo, []string{pb.UserRole_SUPER_ADMIN.String(), pb.UserRole_ADMIN.String()})
	if err != nil {
		return err
	}
	now := time.Now()
	user.CreatedBy = adminID
	user.CreatedAt = now
	user.UpdatedAt = now
	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByID(ctx context.Context, userID primitive.ObjectID) (*models.User, error) {
	user, err := s.userRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetListUser(ctx context.Context, offset, limit int) ([]*models.User, int64, error) {
	if _, err := Authorized(ctx, s.userRepo, []string{
		pb.UserRole_SUPER_ADMIN.String(),
		pb.UserRole_ADMIN.String(),
		pb.UserRole_MANAGER.String(),
	}); err != nil {
		return nil, 0, err
	}
	users, err := s.userRepo.FindAll(ctx, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.userRepo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (s *userService) UpdateUser(ctx context.Context, userID primitive.ObjectID, user *models.User) error {
	user.UpdatedAt = time.Now()
	if err := s.userRepo.Update(ctx, userID, user); err != nil {
		return err
	}
	return nil
}

func (s *userService) DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	if err := s.userRepo.Delete(ctx, userID); err != nil {
		return err
	}
	return nil
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
