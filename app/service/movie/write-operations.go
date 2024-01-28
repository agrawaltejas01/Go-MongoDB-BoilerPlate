package movieService

import (
	"errors"
	movieRepo "shive-app/app/repositories/movie"
	genreService "shive-app/app/service/genre"
	"shive-app/database/models"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMovie(movieObject models.Movie) (primitive.ObjectID, error) {
	// Check if movie already exists
	movieAlreadyExistsError := CheckIfMovieAlreadyExistsByName(movieObject.Name)
	if movieAlreadyExistsError != nil {
		return primitive.NilObjectID, movieAlreadyExistsError
	}

	// Check if genre exists or not
	_, genreFindErr := genreService.GetGenreDataFromGenreId(movieObject.Genre_id)
	if genreFindErr != nil {
		return primitive.NilObjectID, errors.New("genre does not exist in db")
	}

	movieObject.Id = primitive.NewObjectID()

	// Validate struct
	if validationErr := validator.New().Struct(movieObject); validationErr != nil {
		return primitive.NilObjectID, errors.New("error in signup validation find operation ")
	}

	// Save in DB
	return movieRepo.AddMovie(movieObject)
}

func UpdateMovie(movieObject models.Movie) error {
	// Check if already exists
	movieObjectInDb, movieFindErr := GetMovieDataFromMovieId(movieObject.Movie_id)
	if movieFindErr == nil {
		// Movie Object Already exists
		movieObject.Id = movieObjectInDb.Id
	} else {
		movieObject.CreatedAt = time.Now()
		movieObject.Id = primitive.NewObjectID()
	}

	// If passed, Check if passed genre exists
	if movieObject.Genre_id != 0 {
		_, genreFindErr := genreService.GetGenreDataFromGenreId(movieObject.Genre_id)
		if genreFindErr != nil {
			return errors.New("genre does not exist in db")
		}
	} else {
		// Genre is not passed, use genre from db
		movieObject.Genre_id = movieObjectInDb.Genre_id
	}

	// Populate empty fields from DB
	if movieObject.Movie_url == "" && movieFindErr == nil {
		movieObject.Movie_url = movieObjectInDb.Movie_url
	}
	if movieObject.Topic == "" && movieFindErr == nil {
		movieObject.Topic = movieObjectInDb.Topic
	}
	if movieObject.Name == "" && movieFindErr == nil {
		movieObject.Name = movieObjectInDb.Name
	}

	// Validate struct
	if validationErr := validator.New().Struct(movieObject); validationErr != nil {
		return errors.New("error in movie validation operation. All of the required fields not provided to upsert")
	}

	// Save in DB
	return movieRepo.UpsertMovie(movieObject)
}

func DeleteMovie(movieId int) error {
	return movieRepo.DeleteMovie(movieId)
}
