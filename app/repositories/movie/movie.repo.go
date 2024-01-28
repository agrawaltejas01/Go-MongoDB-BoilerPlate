package movieRepo

import (
	"context"
	"shive-app/database"
	"shive-app/database/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var movieModel *mongo.Collection

func init() {
	movieModel, _ = database.GetCollection("movie")
}

func findOne(query interface{}, projection bson.M) (models.Movie, error) {
	result := database.FindOne(movieModel, query, projection)

	var movie models.Movie
	err := result.Decode(&movie)
	return movie, err
}

func buildMoviesSlice(cursor *mongo.Cursor, context context.Context) ([]models.Movie, error) {
	var result []models.Movie

	for cursor.Next(context) {
		var user models.Movie
		cursor.Decode(&user)

		result = append(result, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	defer cursor.Close(context)
	return result, nil
}

func find(query interface{}, projection bson.M) ([]models.Movie, error) {
	cursor, context, err := database.Find(movieModel, query, projection)
	if err != nil {
		return nil, err
	}
	return buildMoviesSlice(cursor, context)
}

func AddMovie(movie models.Movie) (primitive.ObjectID, error) {

	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()
	return database.InsertOne(movieModel, movie)
}

func UpsertMovie(movie models.Movie) error {
	findQuery := bson.M{
		"_id": movie.Id,
	}

	updateQuery := bson.M{
		"$set": bson.M{
			"name":       movie.Name,
			"topic":      movie.Topic,
			"genre_id":   movie.Genre_id,
			"movie_url":  movie.Movie_url,
			"updated_at": time.Now(),
			"movie_id":   movie.Movie_id,
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := database.UpdateOne(movieModel, findQuery, updateQuery, opts)

	return err
}

func DeleteMovie(movieId int) error {
	return database.DeleteOne(movieModel, bson.M{"movie_id": movieId})
}

func FindByName(name string) (models.Movie, error) {
	return findOne(
		bson.M{
			"name": name,
		},
		bson.M{})
}

func FindByMovieId(movieId int) (models.Movie, error) {
	return findOne(
		bson.M{
			"movie_id": movieId,
		},
		bson.M{})
}

func FindByNamePattern(name string) ([]models.Movie, error) {

	filter := bson.D{{Key: "name", Value: primitive.Regex{Pattern: name, Options: "i"}}}

	return find(
		filter,
		bson.M{})
}

func GetAllMoviesForAdmin() ([]models.Movie, error) {
	groupStage := bson.D{{
		Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
			{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
			{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
		},
	}}

	pipeline := mongo.Pipeline{
		groupStage,
	}

	cursor, context, err := database.Aggregate(movieModel, pipeline)

	if err != nil {
		return nil, err
	}

	var allMovies []bson.M
	if err = cursor.All(context, &allMovies); err != nil {
		return nil, err
	}

	var movieSlice []models.Movie

	for _, elem := range allMovies[0]["data"].(primitive.A) {
		if doc, ok := elem.(bson.M); ok { // Check if it's a bson.M document
			var movie models.Movie
			bsonBytes, _ := bson.Marshal(doc)
			if err := bson.Unmarshal(bsonBytes, &movie); err == nil {
				movieSlice = append(movieSlice, movie)
			}
		}
	}

	return movieSlice, nil
}

func FindByGenreId(genreId int) ([]models.Movie, error) {

	filter := bson.D{{Key: "genre_id", Value: genreId}}

	return find(
		filter,
		bson.M{})
}
