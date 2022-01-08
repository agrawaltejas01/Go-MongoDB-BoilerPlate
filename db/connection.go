package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
const uri = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

var client *mongo.Client
var db *mongo.Database
var err error

func connectDB() (*mongo.Database, error) {
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

	db = client.Database("test")

	return db, nil
}

func GetCollection(collection string) (*mongo.Collection, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	return db.Collection(collection), nil
}

func DisconnectDB() {

	if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

}
