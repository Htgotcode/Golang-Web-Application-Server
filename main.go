package main

import (
	"log"
	"os"
	"text/template"

	"github.com/Htgotcode/Golang-Web-Application-Server/database"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
