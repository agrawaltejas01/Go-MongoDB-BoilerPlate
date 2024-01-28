package serverResponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerError(context *gin.Context, errorMsg string) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error":   errorMsg,
	})
}

func BadRequestServerError(context *gin.Context, errorMsg string) {
	context.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   errorMsg,
	})
}

func SuccessResponse(context *gin.Context, data map[string]interface{}, statusCode_optional int) {

	if statusCode_optional == 0 {
		statusCode_optional = http.StatusOK
	}

	data["success"] = true

	// json.Unmarshal([]byte(data), &out)
	// result, _ := jsonmerge.Merge(data, out)

	context.JSON(statusCode_optional, data)
}

func UnauthorizedRequest(context *gin.Context) {
	context.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"error":   "Unauthorized to execute request",
	})
}
