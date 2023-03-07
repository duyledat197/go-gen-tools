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

type userRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(coll *mongo.Collection) repositories.UserRepository {
	return &userRepository{
		coll,
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.User, opts ...repositories.Options) error {
	opt := &options.InsertOneOptions{}
	if _, err := r.coll.InsertOne(ctx, user, opt); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, filter, user *models.User, opts ...repositories.Options) error {
	opt := &options.UpdateOptions{}
	result, err := r.coll.UpdateMany(ctx, filter, primitive.M{
		"set": user,
	}, opt)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("update not effected")
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, filter *models.User, opts ...repositories.Options) error {
	opt := &options.DeleteOptions{}
	result, err := r.coll.DeleteMany(ctx, filter, opt)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("delete not effected")
	}
	return nil
}

func (r *userRepository) GetList(ctx context.Context, filter *models.User, offset, limit int, opts ...repositories.Options) ([]*models.User, error) {
	opt := &options.FindOptions{}
	var result []*models.user
	cur, err := r.coll.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.User, error) {
	opt := &options.FindOneOptions{}
	result := &models.user{}
	if err := r.coll.FindOne(ctx, &models.user{ID: id}, opt).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}
