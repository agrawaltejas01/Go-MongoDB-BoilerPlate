package genreRepo

import (
	"shive-app/database"
	"shive-app/database/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var genreModel *mongo.Collection

func init() {
	genreModel, _ = database.GetCollection("genre")
}

func findOne(query bson.M, projection bson.M) (models.Genre, error) {
	result := database.FindOne(genreModel, query, projection)

	var genre models.Genre
	err := result.Decode(&genre)
	return genre, err
}

func AddGenre(genre models.Genre) (primitive.ObjectID, error) {

	genre.CreatedAt = time.Now()
	genre.UpdatedAt = time.Now()
	return database.InsertOne(genreModel, genre)
}

func UpsertGenre(genre models.Genre) error {
	findQuery := bson.M{
		"_id": genre.Id,
	}

	updateQuery := bson.M{
		"$set": bson.M{
			"name":       genre.Name,
			"genre_id":   genre.Genre_id,
			"updated_at": time.Now(),
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := database.UpdateOne(genreModel, findQuery, updateQuery, opts)

	return err
}

func DeleteGenre(genreId int) error {
	return database.DeleteOne(genreModel, bson.M{"genre_id": genreId})
}

func FindByName(name string) (models.Genre, error) {
	return findOne(
		bson.M{
			"name": name,
		},
		bson.M{})
}

func FindByGenreId(genreId int) (models.Genre, error) {
	return findOne(
		bson.M{
			"genre_id": genreId,
		},
		bson.M{})
}

func GetAllGenresForAdmin() ([]models.Genre, error) {
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

	cursor, context, err := database.Aggregate(genreModel, pipeline)

	if err != nil {
		return nil, err
	}

	var allGenres []bson.M
	if err = cursor.All(context, &allGenres); err != nil {
		return nil, err
	}

	var genreSlice []models.Genre

	for _, elem := range allGenres[0]["data"].(primitive.A) {
		if doc, ok := elem.(bson.M); ok { // Check if it's a bson.M document
			var genre models.Genre
			bsonBytes, _ := bson.Marshal(doc)
			if err := bson.Unmarshal(bsonBytes, &genre); err == nil {
				genreSlice = append(genreSlice, genre)
			}
		}
	}

	return genreSlice, nil
}
