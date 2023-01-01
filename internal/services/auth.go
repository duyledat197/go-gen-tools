package services

import (
	"context"
	"fmt"

	"github.com/duyledat197/interview-hao/internal/models"
	"github.com/duyledat197/interview-hao/internal/repositories"
	"github.com/duyledat197/interview-hao/utils"
	"github.com/duyledat197/interview-hao/utils/helper"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service interface
type AuthService interface {
	Login(ctx context.Context, email, password string) (*models.User, string, error)
	Register(ctx context.Context, user *models.User, password string) error
}

type authService struct {
	userRepo repositories.UserRepository
}

func (s *authService) Login(ctx context.Context, email, password string) (*models.User, string, error) {
	if isEmail := helper.IsEmail(email); !isEmail {
		return nil, "", models.ErrInvalidEmail
	}
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", fmt.Errorf("email not found: %s, %w", email, err)
	}
	if isCorrect := models.IsCorrectPassword(password, user.HashedPassword); !isCorrect {
		return nil, "", models.ErrWrongEmailOrPassword
	}

	tokenString, err := utils.GenerateToken(user.ID.Hex())
	if err != nil {
		return nil, "", err
	}
	return user, tokenString, nil
}

func (s *authService) Register(ctx context.Context, user *models.User, password string) error {

	if isEmail := helper.IsEmail(user.Email); !isEmail {
		return models.ErrInvalidEmail
	}
	if isPassword := helper.IsPassword(password); !isPassword {
		return models.ErrInvalidPassword
	}
	_, err := s.userRepo.FindByEmail(ctx, user.Email)
	if err == nil {
		return models.ErrUserAlreadyExist
	}
	if err != models.ErrUnknowUser {
		return err
	}

	user.HashedPassword, err = models.HashPassword(password)
	user.ID = primitive.NewObjectID()
	if err != nil {
		return err
	}
	err = s.userRepo.Create(ctx, user)
	return err
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}
