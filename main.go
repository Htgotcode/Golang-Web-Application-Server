package main

import (
	"os"
	"text/template"

	"github.com/Htgotcode/Golang-Web-Application-Server/database"
	"github.com/Htgotcode/Golang-Web-Application-Server/routes"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Client = database.DBinstance()

//var accountCollection *mongo.Collection = database.OpenCollection(DB, "account_db", "accounts")
//var validate = validator.New()

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

	r.Run(":" + port)
}
