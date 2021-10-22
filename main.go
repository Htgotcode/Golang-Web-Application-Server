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

	//GETS
	r.GET("/", Handler)
	r.GET("/card-add", Handler)
	r.GET("/all-cards-response", routes.GetCards)
	r.GET("/all-cards", Handler)
	r.GET("/profile", Handler)
	r.GET("/card", routes.GetMarket)
	r.GET("/account", routes.GetAccount)
	//r.GET("/card-listing-name", routes.GetCardListingsByName)
	r.GET("/card-listing", Handler)
  r.GET("/cart-add", Handler)
	r.GET("/cart", Handler)
	r.GET("/cart-response", routes.GetCart)

	//POSTS
	r.POST("/card-add", routes.AddNewcard)
	r.POST("/card-listing-add", routes.CreateCardList)
  
	//POSTS
	r.POST("/card-add", routes.AddNewcard)
	r.POST("/cart-create", routes.CreateCart)
	r.POST("/cart-response/:userid", routes.GetCartByUserId)

	//DELETE
	r.DELETE("/card:id", routes.RemoveCard)
	r.DELETE("/cart-response:id", routes.RemoveCart)

	r.Run(":" + port)
}
