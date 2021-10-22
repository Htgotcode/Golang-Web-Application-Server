package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/Htgotcode/Golang-Web-Application-Server/api/routes"
	"github.com/Htgotcode/Golang-Web-Application-Server/database"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Jwks stores a slice of JSON Web Keys
type Response struct {
	Message string `json:"message"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

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

var jwtMiddleWare *jwtmiddleware.JWTMiddleware

func main() {

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			aud := os.Getenv("AUTH0_API_AUDIENCE")
			checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAudience {
				return token, errors.New("Invalid audience.")
			}
			// verify iss claim
			iss := os.Getenv("AUTH0_DOMAIN")
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			cert, err := getPemCert(token)
			if err != nil {
				log.Fatalf("could not get cert: %+v", err)
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	jwtMiddleWare = jwtMiddleware

	// ... the rest of the code below this function doesn't change yet

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
	r.GET("/", authMiddleware(), Handler)
	r.GET("/card-add", Handler)
	r.GET("/all-cards-response", routes.GetCards)
	r.GET("/all-cards", Handler)
	r.GET("/profile", Handler)
	r.GET("/card", routes.GetMarket)
	r.GET("/account", routes.GetAccount)
	r.GET("/card-listing", Handler)
	r.GET("/cart", Handler)
	r.GET("/cart-response", routes.GetCart)

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

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client secret key
		err := jwtMiddleWare.CheckJWT(c.Writer, c.Request)
		if err != nil {
			// Token not found
			fmt.Println(err)
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Unauthorized"))
			return
		}
	}

}
func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(os.Getenv("AUTH0_DOMAIN") + ".well-known/jwks.json")
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	x5c := jwks.Keys[0].X5c
	for k, v := range x5c {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + v + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		return cert, errors.New("unable to find appropriate key.")
	}

	return cert, nil
}
