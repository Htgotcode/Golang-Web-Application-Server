package main

import (
	"os"
	"text/template"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
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
