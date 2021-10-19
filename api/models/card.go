package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	ID           primitive.ObjectID `bson:"_id"`
	Brand        string             `json:"brand" validate:"required"`
	Set          string             `json:"set" validate:"required"`
	Rarity       string             `json:"rarity" validate:"required"`
	SellingPrice uint               `json:"selling_price" validate:"required"`
	UploadedAt   time.Time          `json:"uploaded_at"`
	CardID       string             `json:"card_id"`
}
