package reviewService

import (
	"errors"
	reviewRepo "shive-app/app/repositories/review"
	movieService "shive-app/app/service/movie"
	"shive-app/database/models"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateReview(reviewObject models.Review) (primitive.ObjectID, error) {

	// Check if movie exists or not
	_, movieFindErr := movieService.GetMovieDataFromMovieId(reviewObject.Movie_id)
	if movieFindErr != nil {
		return primitive.NilObjectID, errors.New("movie does not exist in db")
	}

	reviewObject.Id = primitive.NewObjectID()
	reviewObject.Review_id = reviewObject.Id.Hex()

	// Validate struct
	if validationErr := validator.New().Struct(reviewObject); validationErr != nil {
		return primitive.NilObjectID, errors.New("error in signup validation find operation ")
	}

	// Save in DB
	return reviewRepo.AddReview(reviewObject)
}

func DeleteReview(reviewId string) error {
	return reviewRepo.DeleteReview(reviewId)
}
