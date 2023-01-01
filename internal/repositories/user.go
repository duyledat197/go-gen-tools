package repositories

import (
	"context"

	"github.com/lalaland/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepository ...
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUserID(ctx context.Context, userID primitive.ObjectID) (*models.User, error)
	FindAll(ctx context.Context, offset, limit int) ([]*models.User, error)
	Update(ctx context.Context, userID primitive.ObjectID, user *models.User) error
	Delete(ctx context.Context, userID primitive.ObjectID) error
	Count(ctx context.Context) (int64, error)
}
