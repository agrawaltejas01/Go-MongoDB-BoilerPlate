package userRepo

import (
	"context"
	"shive-app/database"
	"shive-app/database/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userModel *mongo.Collection

func init() {
	userModel, _ = database.GetCollection("users")
}

func insertOne(collection *mongo.Collection, user models.User) (primitive.ObjectID, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(context, user)

	return result.InsertedID.(primitive.ObjectID), err
}

func AddUser(user models.User) (primitive.ObjectID, error) {
	return insertOne(userModel, user)
}
