
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
	db *models.Queries
}

func NewUserRepository(q *models.Queries) UserRepository {
	return &userRepository{
		db: q,
	}
}

func (u *userRepository) Create(ctx context.Context, user *models.User) error {
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
