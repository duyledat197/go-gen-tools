package repositories

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetList(ctx context.Context, offset, limit int) ([]*models.User, error)
	Update(ctx context.Context, id string, user *models.User) error
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db models.DBTX
}

func NewUserRepository(db models.DBTX) UserRepository {
	return &userRepository{
		db,
	}
}

func (u *userRepository) Create(ctx context.Context, user *models.User) error {
	q := models.New(u.db)
	if _, err := q.CreateUser(ctx, models.CreateUserParams{
		ID:     user.ID,
		Name:   user.Name,
		Type:   user.Type,
		TeamID: user.TeamID,
	}); err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Update(ctx context.Context, id string, user *models.User) error {
	return nil
}

func (u *userRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (u *userRepository) GetList(ctx context.Context, offset, limit int) ([]*models.User, error) {
	return nil, nil
}

func (u *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}
