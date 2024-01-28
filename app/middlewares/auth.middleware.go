package authMiddlewares

import (
	authService "shive-app/app/service/auth"
	serverResponse "shive-app/lib/server-response"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("token")

	if token == "" {
		serverResponse.BadRequestServerError(context, "token required")
		context.Abort()
		return
	}

	jwtData, err := authService.ValidateToken(token)

	if err != "" {
		serverResponse.BadRequestServerError(context, "Invalid token - "+err)
		context.Abort()
		return
	}

	context.Set("email", jwtData.Email)
	context.Set("name", jwtData.Name)
	context.Set("userName", jwtData.Username)
	context.Set("userId", jwtData.User_Id)
	context.Set("userType", jwtData.User_type)

	context.Next()

}
