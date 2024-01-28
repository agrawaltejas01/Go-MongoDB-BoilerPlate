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

func GetGenreDataFromGenreId(genreId int) (models.Genre, error) {
	genreInDB, findErr := genreRepo.FindByGenreId(genreId)
	if findErr != nil {
		msg := "Error in finding genre in db"
		if findErr == mongo.ErrNoDocuments {
			msg = "Genre Not found"
		}
		return models.Genre{}, errors.New(msg)
	}

	return genreInDB, nil
}

func GetAllGenres() ([]models.Genre, error) {
	genres, err := genreRepo.GetAllGenresForAdmin()

	if err != nil {
		return nil, errors.New("error in getting genres")
	}

	return genres, nil
}
