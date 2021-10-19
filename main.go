package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/Htgotcode/Golang-Web-Application-Server/database"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

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

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MongoURL")

	client, ctx, cancel, err := database.ConnectDB(MongoDb)
	if err != nil {
		panic(err)
	}

	defer database.CloseConnection(client, ctx, cancel)

	database.CheckConnection(client, ctx)

	accountsCollection := database.OpenCollection(client, "accounts")
	fmt.Println(accountsCollection)
	cursor, err := accountsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var accounts []bson.M
	if err = cursor.All(ctx, &accounts); err != nil {
		log.Fatal(err)
	}
	fmt.Println(accounts)

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(static.Serve("/", static.LocalFile("./ui/build", true)))

	r.Group("/api")

	r.GET("/", Handler)
	r.GET("/card", Handler)
	r.GET("/profile", Handler)
	r.GET("/uploads", Handler)
	r.GET("/market", Handler)

	r.Run(":" + port)
}
