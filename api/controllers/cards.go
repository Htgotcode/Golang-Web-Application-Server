package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Htgotcode/Golang-Web-Application-Server/api/models"
	"github.com/Htgotcode/Golang-Web-Application-Server/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
var card models.Card

//connect to the database and open a card collection
var cardCollection *mongo.Collection = database.OpenCollection(database.Client, "card")

func CreateCards(c *gin.Context) {
	//bind the object that comes in with the declared varaible. thrrow an error if one occurs
	if err := c.BindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// use the validation packge to verify that all items coming in meet the requirements of the struct
	validationErr := validate.Struct(card)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	card.UploadedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	card.ID = primitive.NewObjectID()

	// assign the the auto generated ID to the primary key attribute
	card.CardID = card.ID.Hex()
	card.Bearing = "SomeBearing"
	card.Brand = "SomeBrand"
	card.Rarity = "Rare"
	card.Set = "Pokemon"
	card.SellingPrice = 5

	//insert the newly created object into mongodb
	result, insertErr := cardCollection.InsertOne(ctx, card)
	if insertErr != nil {
		msg := fmt.Sprintf("Card item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()

	//return the id of the created object to the frontend
	c.JSON(http.StatusOK, result)
}
