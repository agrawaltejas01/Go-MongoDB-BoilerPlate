package genreService

import (
	"errors"
	genreRepo "shive-app/app/repositories/genre"
	"shive-app/database/models"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateGenre(genreObject models.Genre) (primitive.ObjectID, error) {
	// Check if already exists
	genreAlreadyExistsError := CheckIfGenreAlreadyExistsByName(genreObject.Name)
	if genreAlreadyExistsError != nil {
		return primitive.NilObjectID, genreAlreadyExistsError
	}

	genreObject.Id = primitive.NewObjectID()

	// Validate struct
	if validationErr := validator.New().Struct(genreObject); validationErr != nil {
		return primitive.NilObjectID, errors.New("error in signup validation find operation ")
	}

	// Save in DB
	return genreRepo.AddGenre(genreObject)
}
