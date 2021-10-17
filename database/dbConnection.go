package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	// Read in connection string
	fmt.Println("Connecting to database...")
	connstring := os.Getenv("cockroachDB_URL")

	// Connect to the "card" database
	db, err := gorm.Open(postgres.Open(connstring), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "card.",
		},
	})
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	return db
}
