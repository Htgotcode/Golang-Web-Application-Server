package main

import (
	"os"
	"text/template"

	"github.com/Htgotcode/Golang-Web-Application-Server/api/middleware"
	"github.com/Htgotcode/Golang-Web-Application-Server/api/routes"
	"github.com/Htgotcode/Golang-Web-Application-Server/database"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Initialise MongoDB database instance
var DB *mongo.Client = database.DBinstance()

// Function to handle parsing files from index.html
func Handler(c *gin.Context) {

	tmpl, err := template.ParseFiles("./ui/build/index.html")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(c.Writer, "")
}

func main() {

	middleware.ValidateKeyGetter()

	// Retrieves the port value stored in the .env file
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialise Gin-gonic middleware instance
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(static.Serve("/", static.LocalFile("./ui/build", true)))

	//GETS
	r.GET("/", middleware.AuthMiddleware(), Handler)
	r.GET("/card-add", Handler)
	r.GET("/all-cards-response", routes.GetCards)
	r.GET("/all-cards", Handler)
	r.GET("/profile", Handler)
	r.GET("/card", routes.GetMarket)
	r.GET("/account", routes.GetAccount)
	r.GET("/card-listing", Handler)
	r.GET("/cart", Handler)
	r.GET("/cart-response", routes.GetCart)
	r.GET("/market", Handler)
	r.GET("/view-card", routes.GetCardListingsByName)

	//POSTS
	r.POST("/card-add", routes.AddNewcard)
	r.POST("/cart-create", routes.CreateCart)
	r.POST("/cart-response/:userid", routes.GetCartByUserId)
	r.POST("/card-listing-add", routes.CreateCardList)

	//DELETE
	r.DELETE("/card:id", routes.RemoveCard)
	r.DELETE("/cart-response:id", routes.RemoveCart)

	r.Run(":" + port)
}

//middleware
