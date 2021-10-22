package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Account database model
type Account struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"createdAt"`
	// Nested object array of purchase history
	PurchaseHistory []struct {
		Name         string    `bson:"name"`
		Description  string    `bson:"description"`
		Brand        string    `bson:"brand" validate:"required"`
		SetName      string    `bson:"set_name" validate:"required"`
		Rarity       string    `bson:"rarity" validate:"required"`
		SellingPrice uint      `bson:"selling_price" validate:"required"`
		UploadedAt   time.Time `bson:"uploaded_at"`
		CardID       string    `bson:"card_id"`
	} `json:"purchaseHistory"`
	// Nested object array of sale history
	SaleHistory []struct {
		Name         string    `bson:"name"`
		Description  string    `bson:"description"`
		Brand        string    `bson:"brand" validate:"required"`
		SetName      string    `bson:"set_name" validate:"required"`
		Rarity       string    `bson:"rarity" validate:"required"`
		SellingPrice uint      `bson:"selling_price" validate:"required"`
		UploadedAt   time.Time `bson:"uploaded_at"`
		CardID       string    `bson:"card_id"`
	} `json:"saleHistory"`
}
