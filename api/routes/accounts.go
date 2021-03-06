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

// Initialise account database collection
var accountCollection *mongo.Collection = database.OpenCollection(database.Client, "account_db", "accounts")

// Function to insert a single new account data item into account collection
func AddAccount(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Account was not created"})
		fmt.Println(insertErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Function to retrieve all account objects in the account collection
func GetAccount(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
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

	fmt.Println(account)

	c.IndentedJSON(http.StatusOK, account)
}

// Function to retrieve accounts according to username parameters
func GetAccountByUsername(c *gin.Context) {

	username := c.Params.ByName("username")

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
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

	fmt.Println(accounts)

	c.JSON(http.StatusOK, accounts)
}

// Function to retrieve specific account by id
func GetAccountById(c *gin.Context) {

	//accountID := c.Params.ByName("_id")

	docID, _ := primitive.ObjectIDFromHex("616d8131ea99fc4e8f9806e1") //For testing without authentication enabled
	fmt.Println(docID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var account bson.M

	if err := accountCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	fmt.Println(account)

	c.JSON(http.StatusOK, account)
}

// Function to update a single account
func UpdateAccount(c *gin.Context) {

	accountID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(accountID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

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

	c.JSON(http.StatusOK, result.ModifiedCount)

}
