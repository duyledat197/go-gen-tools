package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name           string             `bson:"name,omitempty" json:"name,omitempty"`
	Email          string             `bson:"email,omitempty" json:"email,omitempty"`
	Phone          string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Age            int                `bson:"age,omitempty" json:"age,omitempty"`
	HashedPassword string             `bson:"hashedPassword,omitempty" json:"hashedPassword,omitempty"`
	Role           string             `bson:"role,omitempty" json:"role,omitempty"`
	CardID         string             `bson:"card_id,omitempty" json:"card_id,omitempty"`
	Gender         string             `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDay       string             `bson:"birth_day,omitempty" json:"birth_day,omitempty"`
	Address        string             `bson:"address,omitempty" json:"address,omitempty"`
	CreatedBy      primitive.ObjectID `bson:"created_by,omitempty" json:"created_by,omitempty"`
	CreatedAt      time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt      time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt      time.Time          `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// HashPassword : hash password using crypto
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// IsCorrectPassword : check password with passwordhash
func IsCorrectPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ErrUnknowUser ...
var ErrUnknowUser = errors.New("unknown user")

// ErrUserAlreadyExist ...
var ErrUserAlreadyExist = errors.New("user already exist")

// ErrWrongEmailOrPassword ...
var ErrWrongEmailOrPassword = errors.New("wrong email or password")

// ErrInvalidEmail ...
var ErrInvalidEmail = errors.New("invalid email")

// ErrInvalidPassword ...
var ErrInvalidPassword = errors.New("password must in 8 - 32 characters")
