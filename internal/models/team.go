package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string               `bson:"name,omitempty" json:"name,omitempty"`
	Description string               `bson:"description,omitempty" json:"description,omitempty"`
	UserIDs     []primitive.ObjectID `bson:"user_ids,omitempty" json:"user_ids,omitempty"`
	ProjectID   primitive.ObjectID   `bson:"project_id,omitempty" json:"project_id,omitempty"`
	DataIDs     []primitive.ObjectID `bson:"data_ids,omitempty" json:"data_ids,omitempty"`
	CreatedBy   primitive.ObjectID   `bson:"created_by,omitempty" json:"created_by,omitempty"`
	CreatedAt   time.Time            `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time            `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt   time.Time            `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
