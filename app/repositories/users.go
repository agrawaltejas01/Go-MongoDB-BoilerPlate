package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/agrawaltejas01/Go-MongoDB-BoilerPlate/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var usersModel *mongo.Collection

func init() {
	usersModel, _ = db.GetCollection("users")
}

func findOne(collection *mongo.Collection, query bson.M, projection bson.M) (db.User, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.FindOne().SetProjection(projection)

	var user db.User
	err := collection.FindOne(context, query, opts).Decode(&user)

	return user, err
}

func buildUsersSlice(cursor *mongo.Cursor, context context.Context) ([]db.User, error) {
	defer cursor.Close(context)
	var result []db.User

	for cursor.Next(context) {
		var user db.User
		cursor.Decode(&user)

		result = append(result, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func find(collection *mongo.Collection, query bson.M, projection bson.M) ([]db.User, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetProjection(projection)

	cur, err := collection.Find(context, query, opts)
	if err != nil {
		return nil, err
	}
	return buildUsersSlice(cur, context)
}

func insertOne(collection *mongo.Collection, user db.User) (primitive.ObjectID, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(context, user)

	return result.InsertedID.(primitive.ObjectID), err
}

func updateOne(collection *mongo.Collection, filter bson.M, update bson.M) error {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(context, filter, update)

	if result.MatchedCount == 0 {
		panic(errors.New("no docs matched"))
	} else if result.ModifiedCount == 0 {
		panic(errors.New("no docs modified"))
	}

	if err != nil {
		panic(err)
	}

	return nil
}

func deleteOne(collection *mongo.Collection, filter bson.M) error {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(context, filter)

	if result.DeletedCount == 0 {
		panic(errors.New("no docs deleted"))
	}

	if err != nil {
		panic(err)
	}

	return nil
}

func GetUsers() ([]db.User, error) {
	return find(usersModel, bson.M{}, bson.M{})
}

func GetUser(id primitive.ObjectID) (db.User, error) {
	return findOne(usersModel, bson.M{"_id": id}, bson.M{})
}

func AddUser(user db.User) (primitive.ObjectID, error) {
	return insertOne(usersModel, user)
}

func UpdateUser(id primitive.ObjectID) error {
	update := bson.M{
		"$set": bson.M{
			"name": "Tejas",
		},
	}
	return updateOne(usersModel, bson.M{"_id": id}, update)
}

func DeleteUser(id primitive.ObjectID) error {
	return deleteOne(usersModel, bson.M{"_id": id})
}
