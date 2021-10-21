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
	SetName      string             `bson:"setname"`
	Rarity       string             `bson:"rarity"`
	SellingPrice string             `bson:"sellingprice"`
	UploadedAt   time.Time          `bson:"uploadedat"`
	ImageUrl     string             `bson:"imageurl"`
	OwnerID      string             `bson:"ownerid"`
	CardID       string             `bson:"cardid"`
}
