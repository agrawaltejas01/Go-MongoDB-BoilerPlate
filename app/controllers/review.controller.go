package controllers

import (
	reviewService "shive-app/app/service/review"
	"shive-app/database/models"
	serverResponse "shive-app/lib/server-response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func returnReviewData(reviewObject models.Review) map[string]interface{} {
	var userData = make(map[string]interface{})
	userData["name"] = reviewObject.Review
	userData["review_id"] = reviewObject.Review_id
	userData["user_id"] = reviewObject.User_id

	return userData
}

func CreateReview(context *gin.Context) {
	var review models.Review
	jsonBindErr := context.BindJSON(&review)

	userId := context.GetString("userId")
	review.User_id = userId

	if jsonBindErr != nil {
		serverResponse.BadRequestServerError(context, "Bad Data received for review object "+jsonBindErr.Error())
		return
	}

	id, creationErr := reviewService.CreateReview(review)

	if creationErr != nil {
		serverResponse.InternalServerError(context, "Error in Creating Review - "+creationErr.Error())
		return
	}

	var result = make(map[string]interface{})
	result["id"] = id

	serverResponse.SuccessResponse(context, result, 0)
}

func GetReviewsForMovieId(context *gin.Context) {
	movieId := context.Param("movieId")
	movieIdInt, strToIntErr := strconv.Atoi(movieId)

	if strToIntErr != nil {
		serverResponse.BadRequestServerError(context, "Invalid Movie Id passed")
		return
	}

	reviews, err := reviewService.GetReviewDataForMovieId(movieIdInt)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Getting Review Data - "+err.Error())
		return
	}

	var reviewSlice []interface{}
	for _, user := range reviews {
		reviewSlice = append(reviewSlice, returnReviewData(user))
	}

	var result = make(map[string]interface{})
	result["data"] = reviewSlice

	serverResponse.SuccessResponse(context, result, 0)
}

func GetAllUserReviews(context *gin.Context) {

	userId := context.GetString("userId")
	reviews, err := reviewService.GetAllUserReviews(userId)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Getting Review Data - "+err.Error())
		return
	}

	var reviewSlice []interface{}
	for _, user := range reviews {
		reviewSlice = append(reviewSlice, returnReviewData(user))
	}

	var result = make(map[string]interface{})
	result["data"] = reviewSlice

	serverResponse.SuccessResponse(context, result, 0)
}

func DeleteReview(context *gin.Context) {
	reviewId := context.Param("reviewId")

	err := reviewService.DeleteReview(reviewId)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Deleting Review Data - "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	result["success"] = true
	serverResponse.SuccessResponse(context, result, 0)
}
