package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Card database model
type Card struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	Brand        string             `bson:"brand" validate:"required"`
	SetName      string             `bson:"set_name" validate:"required"`
	Rarity       string             `bson:"rarity" validate:"required"`
	SellingPrice float64            `bson:"selling_price" validate:"required"`
	UploadedAt   time.Time          `bson:"uploaded_at"`
	CardID       string             `bson:"card_id"`
	Image        string             `bson:"image_url"`
	OwnerID      string             `bson:"owner_id"`
}
