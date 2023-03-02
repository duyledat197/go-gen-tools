package mongo

import (
	"context"
	"fmt"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type hubRepository struct {
	coll *mongo.Collection
}

func NewHubRepository(coll *mongo.Collection) repositories.HubRepository {
	return &hubRepository{
		coll: coll,
	}
}

func (r *hubRepository) Create(ctx context.Context, hub *models.Hub, opts ...repositories.Options) error {
	if _, err := r.coll.InsertOne(ctx, hub, &options.InsertOneOptions{}); err != nil {
		return err
	}
	return nil
}

func (r *hubRepository) Update(ctx context.Context, filter *models.Hub, hub *models.Hub, opts ...repositories.Options) error {
	result, err := r.coll.UpdateMany(ctx, filter, primitive.M{
		"set": hub,
	}, &options.UpdateOptions{})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("update not effected")
	}
	return nil
}

func (r *hubRepository) Delete(ctx context.Context, id string, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *hubRepository) GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.Hub, error) {
	q := models.New(r.db)
	return nil, nil
}

func (r *hubRepository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.Hub, error) {
	q := models.New(r.db)
	return nil, nil
}
