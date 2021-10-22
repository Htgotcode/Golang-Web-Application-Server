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
	// Retrieves the port value stored in the .env file
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Initialise Gin-gonic middleware instance
	r := gin.New()
	// Gin logger middleware to write logs
	r.Use(gin.Logger())
	// Attaches middleware to the router and returns middleware handler that serves statics files in given directory
	r.Use(static.Serve("/", static.LocalFile("./ui/build", true)))
	// Creates router group API
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

	//POSTS
	r.POST("/card-add", routes.AddNewcard)

	//DELETE
	r.DELETE("/card:id", routes.RemoveCard)

	r.Run(":" + port)
}
