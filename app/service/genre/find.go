package genreService

import (
	"errors"
	genreRepo "shive-app/app/repositories/genre"
	"shive-app/database/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func CheckIfGenreAlreadyExistsByName(name string) error {
	nameGenre, nameGenreFindErr := genreRepo.FindByName(name)

	if nameGenreFindErr != nil && nameGenreFindErr != mongo.ErrNoDocuments {
		return errors.New("error in signup validation find operation ")
	}

	if nameGenre != (models.Genre{}) {
		if nameGenre.Name == name {
			return errors.New(name + " Genre Already exists")
		}
	}

	return nil
}
