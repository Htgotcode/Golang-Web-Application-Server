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

var validate = validator.New()

var accountCollection *mongo.Collection = database.OpenCollection(database.Client, "account_db", "accounts")

func AddAccount(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var account models.Account

	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(account)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	account.ID = primitive.NewObjectID()

	result, insertErr := accountCollection.InsertOne(ctx, account)
	if insertErr != nil {
		msg := fmt.Sprintf("Account was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result)
}

func GetAccount(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var account []bson.M

	cursor, err := accountCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(account)

	c.IndentedJSON(http.StatusOK, account)
}

func GetAccountByUsername(c *gin.Context) {

	username := c.Params.ByName("username")

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var accounts []bson.M

	cursor, err := accountCollection.Find(ctx, bson.M{"username": username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &accounts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(accounts)

	c.JSON(http.StatusOK, accounts)
}

func GetAccountById(c *gin.Context) {

	//accountID := c.Params.ByName("_id")
	//fmt.Println(accountID)
	docID, _ := primitive.ObjectIDFromHex("616d8131ea99fc4e8f9806e1")
	fmt.Println(docID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var account bson.M

	if err := accountCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(account)

	c.JSON(http.StatusOK, account)
}

func UpdateAccount(c *gin.Context) {

	accountID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(accountID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	// type Waiter struct {
	// 	Server *string `json:"server"`
	// }

	var account models.Account

	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	result, err := accountCollection.UpdateOne(ctx, bson.M{"_id": docID},
		bson.D{
			{"$set", bson.D{{"username", account.Username}}},
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

// 	result, err := orderCollection.ReplaceOne(
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
