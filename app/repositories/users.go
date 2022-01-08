package repositories

import (
	"context"
	"time"

	"github.com/agrawaltejas01/Go-MongoDB-BoilerPlate/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersModel *mongo.Collection

func init() {
	usersModel, _ = db.GetCollection("users")
}

func findOne(query bson.M, projection bson.M) (db.User, error) {
	result := db.FindOne(usersModel, query, projection)

	var user db.User
	err := result.Decode(&user)
	return user, err
}

func buildUsersSlice(cursor *mongo.Cursor, context context.Context) ([]db.User, error) {
	var result []db.User

	for cursor.Next(context) {
		var user db.User
		cursor.Decode(&user)

		result = append(result, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	defer cursor.Close(context)
	return result, nil
}

func find(query bson.M, projection bson.M) ([]db.User, error) {
	cursor, context, err := db.Find(usersModel, query, projection)
	if err != nil {
		return nil, err
	}
	return buildUsersSlice(cursor, context)
}

func insertOne(collection *mongo.Collection, user db.User) (primitive.ObjectID, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(context, user)

	return result.InsertedID.(primitive.ObjectID), err
}

func GetUsers() ([]db.User, error) {
	return find(bson.M{}, bson.M{})
}

func GetUser(id primitive.ObjectID) (db.User, error) {
	return findOne(bson.M{"_id": id}, bson.M{})
}

func AddUser(user db.User) (primitive.ObjectID, error) {
	return insertOne(usersModel, user)
}

func UpdateUser(_id primitive.ObjectID) error {
	update := bson.M{
		"$set": bson.M{
			"name": "TSA",
		},
	}
	return db.UpdateOne(usersModel, bson.M{"_id": _id}, update)
}

func DeleteUser(_id primitive.ObjectID) error {
	return db.DeleteOne(usersModel, bson.M{"_id": _id})
}
