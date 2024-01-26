package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var db *mongo.Database
var err error

func ConnectDB() (*mongo.Database, error) {

	// Connection URI
	uri := os.Getenv("DB_URI")
	fmt.Println(uri)

	// Create a new client and connect to the server
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	db = client.Database("shive-movie")

	return db, nil
}

func GetCollection(collection string) (*mongo.Collection, error) {
	return db.Collection(collection), nil
}

func DisconnectDB() {

	if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

}
