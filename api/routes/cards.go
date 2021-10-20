package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	//"api/models"

	"github.com/Htgotcode/Golang-Web-Application-Server/api/models"
	"github.com/Htgotcode/Golang-Web-Application-Server/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validateCard = validator.New()
var cardCollection *mongo.Collection = database.OpenCollection(database.Client, "cards_db", "card_collection")
var marketCollection *mongo.Collection = database.OpenCollection(database.Client, "marketplace_db", "marketplace_collection")
var validate = validator.New()

func AddNewcard(c *gin.Context) {

	var marketplaceCollection *mongo.Collection = database.OpenCollection(database.Client, "marketplace_db", "marketplace_collection")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
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
		msg := fmt.Sprintf("Card item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result)
}

func GetMarket(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

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

// func UpdateWaiter(c *gin.Context) {

// 	orderID := c.Params.ByName("id")
// 	docID, _ := primitive.ObjectIDFromHex(orderID)

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	type Waiter struct {
// 		Server *string `json:"server"`
// 	}

// 	var waiter Waiter

// 	if err := c.BindJSON(&waiter); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	result, err := cardCollection.UpdateOne(ctx, bson.M{"_id": docID},
// 		bson.D{
// 			{"$set", bson.D{{"server", waiter.Server}}},
// 		},
// 	)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	defer cancel()

// 	c.JSON(http.StatusOK, result.ModifiedCount)

// }

//update the order
// func UpdateOrder(c *gin.Context) {

// 	orderID := c.Params.ByName("id")
// 	docID, _ := primitive.ObjectIDFromHex(orderID)

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	var order models.Order

// 	if err := c.BindJSON(&order); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	validationErr := validate.Struct(order)
// 	if validationErr != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
// 		fmt.Println(validationErr)
// 		return
// 	}

// 	result, err := cardCollection.ReplaceOne(
// 		ctx,
// 		bson.M{"_id": docID},
// 		bson.M{
// 			"dish":   order.Dish,
// 			"price":  order.Price,
// 			"server": order.Server,
// 			"table":  order.Table,
// 		},
// 	)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	defer cancel()

// 	c.JSON(http.StatusOK, result.ModifiedCount)
// }

// //delete an order given the id
// func DeleteOrder(c *gin.Context) {

// 	orderID := c.Params.ByName("id")
// 	docID, _ := primitive.ObjectIDFromHex(orderID)

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	result, err := cardCollection.DeleteOne(ctx, bson.M{"_id": docID})

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	defer cancel()

// 	c.JSON(http.StatusOK, result.DeletedCount)
// }
