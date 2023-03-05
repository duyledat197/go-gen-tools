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
	opt := &options.InsertOneOptions{}
	if _, err := r.coll.InsertOne(ctx, hub, opt); err != nil {
		return err
	}
	return nil
}

func (r *hubRepository) Update(ctx context.Context, filter, hub *models.Hub, opts ...repositories.Options) error {
	opt := &options.UpdateOptions{}
	result, err := r.coll.UpdateMany(ctx, filter, primitive.M{
		"set": hub,
	}, opt)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("update not effected")
	}
	return nil
}

func (r *hubRepository) Delete(ctx context.Context, filter *models.Hub, opts ...repositories.Options) error {
	opt := &options.DeleteOptions{}
	result, err := r.coll.DeleteMany(ctx, filter, opt)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("update not effected")
	}
	return nil
}

func (r *hubRepository) GetList(ctx context.Context, filter *models.Hub, offset, limit int, opts ...repositories.Options) ([]*models.Hub, error) {
	opt := &options.FindOptions{}
	var result []*models.Hub
	cur, err := r.coll.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *hubRepository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.Hub, error) {
	opt := &options.FindOneOptions{}
	result := &models.Hub{}
	if err := r.coll.FindOne(ctx, &models.Hub{ID: id}, opt).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}
