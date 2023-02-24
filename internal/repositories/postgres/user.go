
package postgres

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
    "github.com/duyledat197/go-gen-tools/internal/repositories"
)

type userRepository struct {
	db models.DBTX
}

func NewUserRepository(db models.DBTX) repositories.UserRepository {
	return &userRepository{
		db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.User, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *userRepository) Update(ctx context.Context, id string, user *models.User, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *userRepository) GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.User, error) {
	q := models.New(r.db)
	return nil, nil
}

func (r *userRepository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.User, error) {
	q := models.New(r.db)
	return nil, nil
}
