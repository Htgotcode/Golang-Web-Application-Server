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

var validateCart = validator.New()

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "cart_db", "cart_connection")

func CreateCart(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var cart models.Cart

	if err := c.BindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(cart)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	cart.ID = primitive.NewObjectID()

	result, insertErr := cartCollection.InsertOne(ctx, cart)
	if insertErr != nil {
		msg := fmt.Sprintf("Cart was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result)
}

func GetCart(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var cart []bson.M

	cursor, err := cartCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(cart)

	c.IndentedJSON(http.StatusOK, cart)
}

func GetCartById(c *gin.Context) {

	cartID := c.Params.ByName("_id")
	docID, _ := primitive.ObjectIDFromHex(cartID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var cart bson.M

	if err := cartCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(cart)

	c.JSON(http.StatusOK, cart)
}

func UpdateCartOneItem(c *gin.Context) {

	cartID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(cartID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var cart models.Cart

	if err := c.BindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	result, err := cartCollection.UpdateOne(ctx, bson.M{"_id": docID},
		bson.D{
			{"$set", bson.D{{"cards", cart.Cards}}},
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.ModifiedCount)

}

func UpdateCart(c *gin.Context) {

	orderID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(orderID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var cart models.Cart

	if err := c.BindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(cart)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	result, err := cartCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"_id":   cart.ID,
			"cards": cart.Cards,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.ModifiedCount)
}
