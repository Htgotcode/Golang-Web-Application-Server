package database

// import (
// 	"fmt"
// 	"os"

// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDB() *gorm.DB {
// 	// Read in connection string
// 	fmt.Println("Connecting to database...")
// 	connstring := os.Getenv("cockroachDB_URL")

// 	// Connect to the "card" database
// 	// db, err := gorm.Open(postgres.Open(connstring), &gorm.Config{
// 	// 	NamingStrategy: schema.NamingStrategy{
// 	// 		TablePrefix: "card.",
// 	// 	},
// 	// })
// 	// if err != nil {
// 	// 	log.Fatal("error configuring the database: ", err)
// 	// }

// 	// config, err := pgxpool.ParseConfig("cockroachDB_URL" + "sslmode=require&pool_max_conns=40")
// 	// if err != nil {
// 	// 	log.Fatal("error configuring the database: ", err)
// 	// }

// 	// // Create a connection pool to the "bank" database.
// 	// dbpool, err := pgxpool.ConnectConfig(context.Background(), config)
// 	// if err != nil {
// 	// 	log.Fatal("error connecting to the database: ", err)
// 	// }
// 	// defer dbpool.Close()

// 	// return dbpool
// }
