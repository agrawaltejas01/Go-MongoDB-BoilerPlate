package controllers

import (
	genreService "shive-app/app/service/genre"
	"shive-app/database/models"
	serverResponse "shive-app/lib/server-response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func returnGenreData(genreObject models.Genre) map[string]interface{} {
	var userData = make(map[string]interface{})
	userData["name"] = genreObject.Name
	userData["genre_id"] = genreObject.Genre_id

	return userData
}

func CreateGenre(context *gin.Context) {
	var genre models.Genre
	jsonBindErr := context.BindJSON(&genre)

	if jsonBindErr != nil {
		serverResponse.BadRequestServerError(context, "Bad Data received for genre object "+jsonBindErr.Error())
		return
	}

	id, creationErr := genreService.CreateGenre(genre)

	if creationErr != nil {
		serverResponse.InternalServerError(context, "Error in Creating Genre - "+creationErr.Error())
		return
	}

	var result = make(map[string]interface{})
	result["id"] = id

	serverResponse.SuccessResponse(context, result, 0)
}

func GetGenreDetails(context *gin.Context) {
	userId := context.Param("genreId")
	userIdInt, strToIntErr := strconv.Atoi(userId)

	if strToIntErr != nil {
		serverResponse.BadRequestServerError(context, "Invalid Genre Id passed")
		return
	}

	genreObject, err := genreService.GetGenreDataFromGenreId(userIdInt)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Getting Genre Data - "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	result["user"] = returnGenreData(genreObject)
	serverResponse.SuccessResponse(context, result, 0)
}
