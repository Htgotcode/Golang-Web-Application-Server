package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("Golang-Web-ApplicationDB").Collection(collectionName)
	//var collection *mongo.Collection = client.Database("accounts_db").Collection(collectionName)

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
