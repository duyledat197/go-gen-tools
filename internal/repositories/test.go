
package repositories

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
)

type TestRepository interface {
	 Create(ctx context.Context, test *models.Test, opts ...Options) error 
	 GetByID(ctx context.Context, id string, opts ...Options) (*models.Test, error) 
	 GetList(ctx context.Context, offset, limit int, opts ...Options) ([]*models.Test, error) 
	 Update(ctx context.Context, test *models.Test, opts ...Options) error 
	 Delete(ctx context.Context, id string, opts ...Options) error 
}

