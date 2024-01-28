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

func GetAllGenreDetails(context *gin.Context) {
	genres, err := genreService.GetAllGenres()

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Getting Genre Data - "+err.Error())
		return
	}

	var genreSlice []interface{}
	for _, user := range genres {
		genreSlice = append(genreSlice, returnGenreData(user))
	}

	var result = make(map[string]interface{})
	result["data"] = genreSlice

	serverResponse.SuccessResponse(context, result, 0)
}

func GetGenreDetails(context *gin.Context) {
	genreId := context.Param("genreId")
	genreIdInt, strToIntErr := strconv.Atoi(genreId)

	if strToIntErr != nil {
		serverResponse.BadRequestServerError(context, "Invalid Genre Id passed")
		return
	}

	genreObject, err := genreService.GetGenreDataFromGenreId(genreIdInt)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Getting Genre Data - "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	result["genre"] = returnGenreData(genreObject)
	serverResponse.SuccessResponse(context, result, 0)
}

func UpdateGenreDetails(context *gin.Context) {
	var genre models.Genre
	jsonBindErr := context.BindJSON(&genre)

	if jsonBindErr != nil {
		serverResponse.BadRequestServerError(context, "Bad Data received for genre object "+jsonBindErr.Error())
		return
	}

	err := genreService.UpdateGenre(genre)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Updating Genre Data - "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	result["success"] = true
	serverResponse.SuccessResponse(context, result, 0)
}

func DeleteGenre(context *gin.Context) {
	genreId := context.Param("genreId")
	genreIdInt, strToIntErr := strconv.Atoi(genreId)

	if strToIntErr != nil {
		serverResponse.BadRequestServerError(context, "Invalid Genre Id passed")
		return
	}

	err := genreService.DeleteGenre(genreIdInt)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Deleting Genre Data - "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	result["success"] = true
	serverResponse.SuccessResponse(context, result, 0)
}
