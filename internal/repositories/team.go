
package repositories

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
)

type TeamRepository interface {
	 Create(ctx context.Context, team *models.Team, opts ...repositories.Options) error 
	 GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.Team, error) 
	 GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.Team, error) 
	 Update(ctx context.Context, id string, team *models.Team, opts ...repositories.Options) error 
	 Delete(ctx context.Context, id string, opts ...repositories.Options) error 
}

