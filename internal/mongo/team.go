package mongo

import (
	"context"

	"github.com/lalaland/backend/internal/models"
	"github.com/lalaland/backend/internal/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type teamRepository struct {
	coll *mongo.Collection
}

func (r *teamRepository) Create(ctx context.Context, team *models.Team) error {
	result, err := r.coll.InsertOne(ctx, team)
	if err != nil {
		return err
	}
	team.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *teamRepository) Count(ctx context.Context) (int64, error) {
	result, err := r.coll.CountDocuments(ctx, &models.Team{})
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (r *teamRepository) FindAll(ctx context.Context, offset int, limit int) ([]*models.Team, error) {
	var result []*models.Team
	opt := &options.FindOptions{}
	opt.SetSkip(int64(offset))
	opt.SetLimit(int64(limit))
	cur, err := r.coll.Find(ctx, &models.Team{}, opt)
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, err
}

func (r *teamRepository) UpdateUsers(ctx context.Context, filter *models.Team, userIDs []primitive.ObjectID) error {
	if _, err := r.coll.UpdateOne(ctx, filter, bson.M{
		"$push": bson.M{
			"user_ids": bson.M{
				"$each": userIDs,
			},
		},
	}); err != nil {
		return err
	}
	return nil
}

func (r *teamRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Team, error) {
	var result models.Team

	if err := r.coll.FindOne(ctx, &models.Team{ID: id}).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *teamRepository) FindByFilter(ctx context.Context, filter *models.Team, offset, limit int) ([]*models.Team, error) {
	var result []*models.Team
	opt := &options.FindOptions{}
	opt.SetSkip(int64(offset))
	opt.SetLimit(int64(limit))
	cur, err := r.coll.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, err
}

func (r *teamRepository) Update(ctx context.Context, team *models.Team) error {
	if _, err := r.coll.UpdateByID(ctx, team.ID, team); err != nil {
		return err
	}
	return nil
}

func NewTeamRepository(coll *mongo.Collection) repositories.TeamRepository {
	// opt := &options.IndexOptions{}
	// opt.SetUnique(true)
	// indexModel := mongo.IndexModel{
	// 	Keys:    bson.D{primitive.E{Key: "email", Value: 1}},
	// 	Options: opt,
	// }

	// if _, err := coll.Indexes().CreateOne(context.TODO(), indexModel); err != nil {
	// 	log.Fatal(err)
	// }
	return &teamRepository{
		coll: coll,
	}
}
