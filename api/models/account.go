package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`

	PurchaseHistory []struct {
		Name         string    `bson:"name"`
		Description  string    `bson:"description"`
		Brand        string    `json:"brand" validate:"required"`
		SetName      string    `json:"set_name" validate:"required"`
		Rarity       string    `json:"rarity" validate:"required"`
		SellingPrice uint      `json:"selling_price" validate:"required"`
		UploadedAt   time.Time `json:"uploaded_at"`
		CardID       string    `json:"card_id"`
	} `json:"purchaseHistory"`

	SaleHistory []struct {
		Name         string    `bson:"name"`
		Description  string    `bson:"description"`
		Brand        string    `json:"brand" validate:"required"`
		SetName      string    `json:"set_name" validate:"required"`
		Rarity       string    `json:"rarity" validate:"required"`
		SellingPrice uint      `json:"selling_price" validate:"required"`
		UploadedAt   time.Time `json:"uploaded_at"`
		CardID       string    `json:"card_id"`
	} `json:"saleHistory"`
}
