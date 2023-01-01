package repositories

import (
	"context"

	"github.com/duyledat197/interview-hao/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamRepository interface {
	Create(ctx context.Context, team *models.Team) error
	Count(ctx context.Context) (int64, error)
	FindAll(ctx context.Context, offset, limit int) ([]*models.Team, error)
	UpdateUsers(ctx context.Context, filter *models.Team, userIDs []primitive.ObjectID) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Team, error)
	FindByFilter(ctx context.Context, filter *models.Team, offset, limit int) ([]*models.Team, error)
	Update(ctx context.Context, team *models.Team) error
}
