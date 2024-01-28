package controllers

import (
	genreService "shive-app/app/service/genre"
	"shive-app/database/models"
	serverResponse "shive-app/lib/server-response"

	"github.com/gin-gonic/gin"
)

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
