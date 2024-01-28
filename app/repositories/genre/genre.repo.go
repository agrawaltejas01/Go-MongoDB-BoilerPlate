package genreRepo

import (
	"shive-app/database"
	"shive-app/database/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
