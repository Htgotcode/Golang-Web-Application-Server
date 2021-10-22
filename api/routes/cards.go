package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Htgotcode/Golang-Web-Application-Server/api/models"
	"github.com/Htgotcode/Golang-Web-Application-Server/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cardCollection *mongo.Collection = database.OpenCollection(database.Client, "cards_db", "card_collection")
var marketCollection *mongo.Collection = database.OpenCollection(database.Client, "marketplace_db", "marketplace_collection")
var validate = validator.New()

//Card functions
func AddNewcard(c *gin.Context) {

	var marketplaceCollection *mongo.Collection = database.OpenCollection(database.Client, "marketplace_db", "marketplace_collection")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var card models.Card

	card.ID = primitive.NewObjectID()
	card.CardID = card.ID.Hex()
	card.UploadedAt = time.Now()

	if err := c.BindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(card)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	result, insertErr := marketplaceCollection.InsertOne(ctx, card)

	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Card item was not created"})
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result)
}

func GetMarket(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var cards []bson.M

	cursor, err := marketCollection.Find(ctx, bson.M{})

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

	c.IndentedJSON(http.StatusOK, cards)
}

func GetCards(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var cards []bson.M

	cursor, err := cardCollection.Find(ctx, bson.M{})

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

	c.IndentedJSON(http.StatusOK, cards)
}

func GetCardsByBrand(c *gin.Context) {

	brand := c.Params.ByName("brand")

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var cards []bson.M

	cursor, err := cardCollection.Find(ctx, bson.M{"brand": brand})
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

func GetCardById(c *gin.Context) {

	cardID := c.Params.ByName("_id")
	docID, _ := primitive.ObjectIDFromHex(cardID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var card bson.M

	if err := cardCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&card); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(card)

	c.JSON(http.StatusOK, card)
}

func GetCardByName(c *gin.Context) {

	name := c.Params.ByName("name")

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var card bson.M

	if err := cardCollection.FindOne(ctx, bson.M{"name": name}).Decode(&card); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(card)

	c.JSON(http.StatusOK, card)
}

func GetPokemonCards(c *gin.Context) {

	cardID := c.Params.ByName("_id")
	docID, _ := primitive.ObjectIDFromHex(cardID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var card bson.M

	if err := cardCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&card); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(card)

	c.JSON(http.StatusOK, card)
}

func RemoveCard(c *gin.Context) {
	cardID := c.Param("id")
	docID, _ := primitive.ObjectIDFromHex(cardID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	result, err := marketCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.IndentedJSON(http.StatusOK, result.DeletedCount)
}
