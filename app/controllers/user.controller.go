package controllers

import (
	userService "shive-app/app/service/user"
	serverResponse "shive-app/lib/server-response"

	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {

	id, creationErr := userService.Create()

	if creationErr != nil {
		serverResponse.InternalServerError(context, "Error in Creating User")
		return
	}

	var result = make(map[string]interface{})
	result["id"] = id

	serverResponse.SuccessResponse(context, result, 0)

}
