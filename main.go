package main

import (
	"os"
	"text/template"

	"github.com/Htgotcode/Golang-Web-Application-Server/api/routes"
	"github.com/Htgotcode/Golang-Web-Application-Server/database"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
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

	//GETS
	r.GET("/", Handler)
	r.GET("/card-add", Handler)
	r.GET("/all-cards-response", routes.GetCards)
	r.GET("/all-cards", Handler)
	r.GET("/profile", Handler)
	r.GET("/card", routes.GetMarket)
	r.GET("/cart", Handler)
	r.GET("/account", routes.GetAccount)
	//r.GET("/card-listing-name", routes.GetCardListingsByName)
	r.GET("/card-listing", Handler)

	//POSTS
	r.POST("/card-add", routes.AddNewcard)
	r.POST("/card-listing-add", routes.CreateCardList)

	//DELETE
	r.DELETE("/card:id", routes.RemoveCard)

	r.Run(":" + port)
}
