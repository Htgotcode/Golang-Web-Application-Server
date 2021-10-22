package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	Brand        string             `bson:"brand"`
	SetName      string             `bson:"set_name"`
	Rarity       string             `bson:"rarity"`
	SellingPrice float64            `bson:"selling_price"`
	UploadedAt   time.Time          `bson:"uploaded_at"`
	CardID       string             `bson:"card_id"`
	Image        string             `bson:"image_url"`
	OwnerID      string             `bson:"owner_id"`
}
