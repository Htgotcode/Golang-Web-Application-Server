package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"createdAt"`

	PurchaseHistory []struct {
		Name         string    `bson:"name"`
		Description  string    `bson:"description"`
		Brand        string    `bson:"brand" validate:"required"`
		SetName      string    `bson:"set_name" validate:"required"`
		Rarity       string    `bson:"rarity" validate:"required"`
		SellingPrice int64     `bson:"selling_price" validate:"required"`
		UploadedAt   time.Time `bson:"uploaded_at"`
		CardID       string    `bson:"card_id"`
	} `json:"purchaseHistory"`

	SaleHistory []struct {
		Name         string    `bson:"name"`
		Description  string    `bson:"description"`
		Brand        string    `bson:"brand" validate:"required"`
		SetName      string    `bson:"set_name" validate:"required"`
		Rarity       string    `bson:"rarity" validate:"required"`
		SellingPrice int64     `bson:"selling_price" validate:"required"`
		UploadedAt   time.Time `bson:"uploaded_at"`
		CardID       string    `bson:"card_id"`
	} `json:"saleHistory"`
}
