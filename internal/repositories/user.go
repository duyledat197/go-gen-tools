package repositories

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User, opts ...Options) error
	GetByID(ctx context.Context, id string, opts ...Options) (*models.User, error)
	GetList(ctx context.Context, offset, limit int, opts ...Options) ([]*models.User, error)
	Update(ctx context.Context, id string, user *models.User, opts ...Options) error
	Delete(ctx context.Context, id string, opts ...Options) error
}
