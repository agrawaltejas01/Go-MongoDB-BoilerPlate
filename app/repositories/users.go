package repositories

import (
	"context"
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

func findOneAndProject(collection *mongo.Collection, query bson.M, projection bson.M) (db.User, error) {
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

func GetUsers() ([]db.User, error) {
	return find(usersModel, bson.M{}, bson.M{})
}

func GetUser(id primitive.ObjectID) (db.User, error) {
	return findOneAndProject(usersModel, bson.M{"_id": id}, bson.M{})
}
