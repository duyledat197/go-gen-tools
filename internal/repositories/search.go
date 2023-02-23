
package repositories

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
)

type SearchRepository interface {
	 Create(ctx context.Context, search *models.Search, opts ...repositories.Options) error 
	 GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.Search, error) 
	 GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.Search, error) 
	 Update(ctx context.Context, id string, search *models.Search, opts ...repositories.Options) error 
	 Delete(ctx context.Context, id string, opts ...repositories.Options) error 
}

