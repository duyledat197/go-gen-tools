
package repositories

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
)

type HubRepository interface {
	 Create(ctx context.Context, hub *models.Hub, opts ...repositories.Options) error 
	 GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.Hub, error) 
	 GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.Hub, error) 
	 Update(ctx context.Context, id string, hub *models.Hub, opts ...repositories.Options) error 
	 Delete(ctx context.Context, id string, opts ...repositories.Options) error 
}

