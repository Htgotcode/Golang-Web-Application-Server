package main

import (
	"os"

	"github.com/Htgotcode/Golang-Web-Application-Server/api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()
	r.Use(gin.Logger())

	r.POST("/cards-create", controllers.CreateCards)

	r.Run(":" + port)
}
