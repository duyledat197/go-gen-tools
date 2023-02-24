
package mongo

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
    "github.com/duyledat197/go-gen-tools/internal/repositories"
)

type searchRepository struct {
	db models.DBTX
}

func NewSearchRepository(db models.DBTX) repositories.SearchRepository {
	return &searchRepository{
		db,
	}
}

func (r *searchRepository) Create(ctx context.Context, search *models.Search, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *searchRepository) Update(ctx context.Context, id string, search *models.Search, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *searchRepository) Delete(ctx context.Context, id string, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *searchRepository) GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.Search, error) {
	q := models.New(r.db)
	return nil, nil
}

func (r *searchRepository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.Search, error) {
	q := models.New(r.db)
	return nil, nil
}
