package movieService

import (
	"errors"
	movieRepo "shive-app/app/repositories/movie"
	"shive-app/database/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func CheckIfMovieAlreadyExistsByName(name string) error {
	nameMovie, nameMovieFindErr := movieRepo.FindByName(name)

	if nameMovieFindErr != nil && nameMovieFindErr != mongo.ErrNoDocuments {
		return errors.New("error in signup validation find operation ")
	}

	if nameMovie != (models.Movie{}) {
		if nameMovie.Name == name {
			return errors.New(name + " Movie Already exists")
		}
	}

	return nil
}

func GetMovieDataFromMovieId(movieId int) (models.Movie, error) {
	movieInDB, findErr := movieRepo.FindByMovieId(movieId)
	if findErr != nil {
		msg := "Error in finding movie in db"
		if findErr == mongo.ErrNoDocuments {
			msg = "Movie Not found"
		}
		return models.Movie{}, errors.New(msg)
	}

	return movieInDB, nil
}

func GetAllMovies() ([]models.Movie, error) {
	movies, err := movieRepo.GetAllMoviesForAdmin()

	if err != nil {
		return nil, errors.New("error in getting movies")
	}

	return movies, nil
}

func SearchMoviesByNamePattern(name string) ([]models.Movie, error) {
	movieInDB, findErr := movieRepo.FindByNamePattern(name)
	if findErr != nil {
		msg := "Error in finding movie in db"
		if findErr == mongo.ErrNoDocuments {
			msg = "Movie Not found"
		}
		return []models.Movie{}, errors.New(msg)
	}

	return movieInDB, nil
}

func SearchMovieByGenreId(genreId int) ([]models.Movie, error) {
	movieInDB, findErr := movieRepo.FindByGenreId(genreId)
	if findErr != nil {
		msg := "Error in finding movie in db"
		if findErr == mongo.ErrNoDocuments {
			msg = "Movie Not found"
		}
		return []models.Movie{}, errors.New(msg)
	}

	return movieInDB, nil
}
