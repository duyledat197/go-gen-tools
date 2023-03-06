
package repositories

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
)

type HubRepository interface {
	 Create(ctx context.Context, hub *models.Hub, opts ...Options) error 
	 GetByID(ctx context.Context, id string, opts ...Options) (*models.Hub, error) 
	 GetList(ctx context.Context, offset, limit int, opts ...Options) ([]*models.Hub, error) 
	 Update(ctx context.Context, hub *models.Hub, opts ...Options) error 
	 Delete(ctx context.Context, id string, opts ...Options) error 
}

