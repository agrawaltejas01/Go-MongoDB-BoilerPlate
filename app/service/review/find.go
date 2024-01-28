package reviewService

import (
	"errors"
	reviewRepo "shive-app/app/repositories/review"
	"shive-app/database/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetReviewDataForMovieId(movieId int) ([]models.Review, error) {
	reviewInDB, findErr := reviewRepo.FindByMovieId(movieId)
	if findErr != nil {
		msg := "Error in finding review in db"
		if findErr == mongo.ErrNoDocuments {
			msg = "Review Not found"
		}
		return []models.Review{}, errors.New(msg)
	}

	return reviewInDB, nil
}

func GetAllUserReviews(userId string) ([]models.Review, error) {
	reviews, err := reviewRepo.FindByUserId(userId)

	if err != nil {
		return nil, errors.New("error in getting reviews")
	}

	return reviews, nil
}
