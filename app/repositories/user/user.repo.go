package userRepo

import (
	"context"
	"shive-app/database"
	"shive-app/database/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userModel *mongo.Collection

func init() {
	userModel, _ = database.GetCollection("users")
}

func findOne(query bson.M, projection bson.M) (models.User, error) {
	result := database.FindOne(userModel, query, projection)

	var user models.User
	err := result.Decode(&user)
	return user, err
}

func insertOne(collection *mongo.Collection, user models.User) (primitive.ObjectID, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	result, err := collection.InsertOne(context, user)

	return result.InsertedID.(primitive.ObjectID), err
}

func AddUser(user models.User) (primitive.ObjectID, error) {
	return insertOne(userModel, user)
}

func FindByEmailOrUserName(email string, userName string) (models.User, error) {
	return findOne(
		bson.M{
			"$or": []interface{}{
				bson.M{"email": email},
				bson.M{"userName": userName},
			},
		},
		bson.M{})
}
