package database

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOne(collection *mongo.Collection, document interface{}) (primitive.ObjectID, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(context, document)

	return result.InsertedID.(primitive.ObjectID), err
}

func FindOne(collection *mongo.Collection, query bson.M, projection bson.M) *mongo.SingleResult {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.FindOne().SetProjection(projection)

	result := collection.FindOne(context, query, opts)

	return result
}

func Find(collection *mongo.Collection, query bson.M, projection bson.M) (*mongo.Cursor, context.Context, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetProjection(projection)

	cur, err := collection.Find(context, query, opts)
	if err != nil {
		return nil, nil, err
	}
	return cur, context, nil
}

func UpdateOne(collection *mongo.Collection, filter bson.M, update bson.M) error {
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

func UpdateMany(collection *mongo.Collection, filter bson.M, update bson.M) error {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateMany(context, filter, update)

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

func DeleteOne(collection *mongo.Collection, filter bson.M) error {
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

func DeleteMany(collection *mongo.Collection, filter bson.M) error {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteMany(context, filter)

	if result.DeletedCount == 0 {
		panic(errors.New("no docs deleted"))
	}

	if err != nil {
		panic(err)
	}

	return nil
}

func Aggregate(collection *mongo.Collection, Pipeline mongo.Pipeline) (*mongo.Cursor, context.Context, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Aggregate(context, Pipeline)

	if err != nil {
		return nil, nil, err
	}

	return cursor, context, nil
}
