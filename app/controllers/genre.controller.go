package controllers

import (
	serverResponse "shive-app/lib/server-response"

	"github.com/gin-gonic/gin"
)

func CreateGenre(context *gin.Context) {
	var result = make(map[string]interface{})
	result["data"] = "123"
	serverResponse.SuccessResponse(context, result, 0)
}
