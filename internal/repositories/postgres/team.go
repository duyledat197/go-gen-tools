
package postgres

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
    "github.com/duyledat197/go-gen-tools/internal/repositories"
)

type teamRepository struct {
	db models.DBTX
}

func NewTeamRepository(db models.DBTX) repositories.TeamRepository {
	return &teamRepository{
		db,
	}
}

func (r *teamRepository) Create(ctx context.Context, team *models.Team, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *teamRepository) Update(ctx context.Context, id string, team *models.Team, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *teamRepository) Delete(ctx context.Context, id string, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *teamRepository) GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.Team, error) {
	q := models.New(r.db)
	return nil, nil
}

func (r *teamRepository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.Team, error) {
	q := models.New(r.db)
	return nil, nil
}
