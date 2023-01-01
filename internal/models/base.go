package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	CreatedBy primitive.ObjectID `bson:"created_by,omitempty" json:"created_by,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
