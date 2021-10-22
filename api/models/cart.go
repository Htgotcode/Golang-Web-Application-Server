package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID    primitive.ObjectID `bson:"_id"`
	Email string             `bson:"email"`
	Total float64            `bson:"total"`

	Cards []Card `bson:"cards"`
}
