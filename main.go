package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/Htgotcode/Golang-Web-Application-Server/api/models"
	"github.com/Htgotcode/Golang-Web-Application-Server/database"
	"github.com/Htgotcode/Golang-Web-Application-Server/routes"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Client = database.DBinstance()

func Handler(c *gin.Context) {
	tmpl, err := template.ParseFiles("./ui/build/index.html")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(c.Writer, "")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(static.Serve("/", static.LocalFile("./ui/build", true)))

	r.Group("/api")

	r.GET("/", Handler)
	r.GET("/card", Handler)
	r.GET("/profile", Handler)
	r.GET("/uploads", Handler)
	r.GET("/market", Handler)

	r.GET("/account", routes.GetAccount)
	r.GET("/card/getCards", routes.GetCards)

	r.POST("/add-card", addNewcard)

	r.Run(":" + port)
}

var validate = validator.New()

func addNewcard(c *gin.Context, name, description, brand, setName, rarity, imageURL, ownerID string, sellingPrice uint) {
	var marketplaceCollection *mongo.Collection = database.OpenCollection(database.Client, "marketplace_db", "marketplace_collection")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var card models.Card

	if err := c.BindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(card)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	card.ID = primitive.NewObjectID()

	card.CardID = card.ID.Hex()
	card.Name = name
	card.Description = description
	card.Brand = brand
	card.SetName = setName
	card.Rarity = rarity
	card.SellingPrice = sellingPrice
	card.UploadedAt = time.Now()
	card.Image = imageURL
	card.OwnerID = ownerID

	result, insertErr := marketplaceCollection.InsertOne(ctx, card)
	if insertErr != nil {
		msg := fmt.Sprintf("Card item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result)

}
