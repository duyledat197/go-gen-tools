
package repositories

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
)

type UserRepository interface {
	 Create(ctx context.Context, user *models.User, opts ...repositories.Options) error 
	 GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.User, error) 
	 GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.User, error) 
	 Update(ctx context.Context, id string, user *models.User, opts ...repositories.Options) error 
	 Delete(ctx context.Context, id string, opts ...repositories.Options) error 
}

