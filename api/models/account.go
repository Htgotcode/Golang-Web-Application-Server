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
		Brand        string    `bson:"brand"`
		SetName      string    `bson:"set_name"`
		Rarity       string    `bson:"rarity"`
		SellingPrice float64   `bson:"sellingprice"`
		UploadedAt   time.Time `bson:"uploaded_at"`
		CardID       string    `bson:"card_id"`
	} `json:"purchaseHistory"`

	SaleHistory []struct {
		Name         string    `bson:"name"`
		Description  string    `bson:"description"`
		Brand        string    `bson:"brand"`
		SetName      string    `bson:"set_name"`
		Rarity       string    `bson:"rarity"`
		SellingPrice float64   `bson:"sellingprice"`
		UploadedAt   time.Time `bson:"uploaded_at"`
		CardID       string    `bson:"card_id"`
	} `json:"saleHistory"`
}
