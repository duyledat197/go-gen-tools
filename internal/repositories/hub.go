package repositories

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
)

type HubRepository interface {
	Create(ctx context.Context, hub *models.Hub) error
	GetByID(ctx context.Context, id string) (*models.Hub, error)
	GetList(ctx context.Context, offset, limit int) ([]*models.Hub, error)
	Update(ctx context.Context, id string, hub *models.Hub) error
	Delete(ctx context.Context, id string) error
}

type hubRepository struct {
	db *models.Queries
}

func NewHubRepository(q *models.Queries) HubRepository {
	return &hubRepository{
		db: q,
	}
}

func (u *hubRepository) Create(ctx context.Context, hub *models.Hub) error {
	if _, err := u.db.CreateHub(ctx, models.CreateHubParams{
		ID:         hub.ID,
		Name:       hub.Name,
		LocationID: hub.LocationID,
	}); err != nil {
		return err
	}
	return nil
}

func (u *hubRepository) Update(ctx context.Context, id string, hub *models.Hub) error {
	return nil
}

func (u *hubRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (u *hubRepository) GetList(ctx context.Context, offset, limit int) ([]*models.Hub, error) {
	return nil, nil
}

func (u *hubRepository) GetByID(ctx context.Context, id string) (*models.Hub, error) {
	return nil, nil
}
