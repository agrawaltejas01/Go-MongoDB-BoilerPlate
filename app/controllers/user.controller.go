package controllers

import (
	userService "shive-app/app/service/user"
	"shive-app/database/models"
	serverResponse "shive-app/lib/server-response"

	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {

	var user models.User
	jsonBindErr := context.BindJSON(&user)

	if jsonBindErr != nil {
		serverResponse.BadRequestServerError(context, "Bad Data received for user object "+jsonBindErr.Error())
		return
	}

	id, creationErr := userService.Create(user)

	if creationErr != nil {
		serverResponse.InternalServerError(context, "Error in Creating User - "+creationErr.Error())
		return
	}

	var result = make(map[string]interface{})
	result["id"] = id

	serverResponse.SuccessResponse(context, result, 0)

}
