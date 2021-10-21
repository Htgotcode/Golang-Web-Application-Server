package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID primitive.ObjectID `bson:"_id"`

	Cards []Card `bson:"cards"`
}
