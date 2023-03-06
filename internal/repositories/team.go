package repositories

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
)

type TeamRepository interface {
	Create(ctx context.Context, team *models.Team, opts ...Options) error
	GetByID(ctx context.Context, id string, opts ...Options) (*models.Team, error)
	GetList(ctx context.Context, offset, limit int, opts ...Options) ([]*models.Team, error)
	Update(ctx context.Context, id string, team *models.Team, opts ...Options) error
	Delete(ctx context.Context, id string, opts ...Options) error
}
