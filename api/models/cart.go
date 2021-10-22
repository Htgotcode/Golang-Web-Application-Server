package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Cart database model
type Cart struct {
	ID primitive.ObjectID `bson:"_id"`

	Cards []Card `bson:"cards"`
}
