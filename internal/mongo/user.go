package mongo

import (
	"context"
	"log"

	"github.com/lalaland/backend/internal/models"
	"github.com/lalaland/backend/internal/repositories"
	"github.com/lalaland/backend/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	coll *mongo.Collection
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	if _, err := r.coll.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var result models.User
	err := r.coll.FindOne(ctx, &models.User{Email: email}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) FindByUserID(ctx context.Context, userID primitive.ObjectID) (*models.User, error) {
	var result models.User
	if err := r.coll.FindOne(ctx, &models.User{ID: userID}).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) FindAll(ctx context.Context, offset, limit int) ([]*models.User, error) {
	var result []*models.User
	opt := &options.FindOptions{}
	opt.SetSkip(int64(offset))
	opt.SetLimit(int64(limit))
	cur, err := r.coll.Find(ctx, bson.M{
		"role": bson.M{
			"$ne": pb.UserRole_SUPER_ADMIN.String(),
		},
	}, opt)
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, err
}

func (r *userRepository) Update(ctx context.Context, userID primitive.ObjectID, user *models.User) error {
	if _, err := r.coll.UpdateByID(ctx, userID, user); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, userID primitive.ObjectID) error {
	return nil
}

func (r *userRepository) Count(ctx context.Context) (int64, error) {
	result, err := r.coll.CountDocuments(ctx, bson.M{
		"role": bson.M{
			"$ne": pb.UserRole_SUPER_ADMIN.String(),
		},
	})
	if err != nil {
		return 0, err
	}
	return result, nil
}

// NewUserRepository ...
func NewUserRepository(coll *mongo.Collection) repositories.UserRepository {
	opt := &options.IndexOptions{}
	opt.SetUnique(true)
	indexModel := mongo.IndexModel{
		Keys:    bson.D{primitive.E{Key: "email", Value: 1}},
		Options: opt,
	}

	if _, err := coll.Indexes().CreateOne(context.TODO(), indexModel); err != nil {
		log.Fatal(err)
	}
	return &userRepository{
		coll: coll,
	}
}
