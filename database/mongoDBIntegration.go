package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	MongoDb := os.Getenv("MongoUrl")

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

//Client Database instance
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

func InsertSingleDoc(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	collection := client.Database(dataBase).Collection(col)

	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func InsertManyDocs(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {

	collection := client.Database(dataBase).Collection(col)

	result, err := collection.InsertMany(ctx, docs)
	return result, err
}

func DBquery(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
	collection := client.Database(dataBase).Collection(col)

	result, err = collection.Find(ctx, query,
		options.Find().SetProjection(field))
	return result, err
}

func UpdateSingleDoc(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

	collection := client.Database(dataBase).Collection(col)

	result, err = collection.UpdateOne(ctx, filter, update)
	return result, err
}

func UpdateManyDocs(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

	collection := client.Database(dataBase).Collection(col)

	result, err = collection.UpdateMany(ctx, filter, update)
	return result, err
}

func DeleteSingleDoc(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {

	collection := client.Database(dataBase).Collection(col)

	result, err = collection.DeleteOne(ctx, query)
	return result, err
}

func DeleteManyDocs(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {

	collection := client.Database(dataBase).Collection(col)

	result, err = collection.DeleteMany(ctx, query)
	return result, err
}

// func CreateAccount(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	var account models.Account
// 	_ = json.NewDecoder(r.Body).Decode(&account)

// 	registerAccount(account)
// 	json.NewEncoder(w).Encode(account)
// }

// func DeleteAccount(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	params := mux.Vars(r)
// 	deleteOneAccount(params["id"])
// 	json.NewEncoder(w).Encode(params["id"])
// }

func GetAllAccountsRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := GetAllAccounts()
	json.NewEncoder(w).Encode(payload)
}

func GetAllAccounts() []primitive.M {
	//accountsCollection := client.Database("account_db").Collection("accounts")

	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// func registerAccount(account models.Account) {
// 	//accountsCollection := client.Database("account_db").Collection("accounts")
// 	insertResult, err := collection.InsertOne(context.Background(), account)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
// }

// func deleteOneAccount(account string) {
// 	//accountsCollection := client.Database("account_db").Collection("accounts")
// 	fmt.Println(account)
// 	id, _ := primitive.ObjectIDFromHex(account)
// 	filter := bson.M{"_id": id}
// 	d, err := collection.DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Deleted Document", d.DeletedCount)
// }
