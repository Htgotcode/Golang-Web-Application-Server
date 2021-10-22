package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Htgotcode/Golang-Web-Application-Server/api/models"
	"github.com/Htgotcode/Golang-Web-Application-Server/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cardListCollection *mongo.Collection = database.OpenCollection(database.Client, "card_list_db", "card_list_connection")

func CreateCardList(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var card models.Card

	card.ID = primitive.NewObjectID()
	card.CardID = card.ID.Hex()
	card.UploadedAt = time.Now()

	if err := c.BindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(card)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	card.ID = primitive.NewObjectID()

	result, insertErr := cardListCollection.InsertOne(ctx, card)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Card listing was not created"})
		fmt.Println(insertErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetCardListingsByName(c *gin.Context) {

	name := c.Params.ByName("name")

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var cards []bson.M

	cursor, err := cardListCollection.Find(ctx, bson.M{"name": name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &cards); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(cards)

	c.JSON(http.StatusOK, cards)
}
