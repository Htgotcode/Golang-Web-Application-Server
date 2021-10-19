package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const connectionString = "Connection String"

// Database Name
const accountDBName = "accountDB"

// Collection name
const accountCollectionName = "accounts"

// collection object/instance
var collection *mongo.Collection

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MongoURL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database(databaseName).Collection(collectionName)

	return collection
}

func ConnectDB(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func CheckConnection(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("MongoDB Connected.")
	return nil
}

func CloseConnection(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
