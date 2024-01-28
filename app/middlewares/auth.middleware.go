package authMiddlewares

import (
	authService "shive-app/app/service/auth"
	serverResponse "shive-app/lib/server-response"

	"github.com/gin-gonic/gin"
)

func authenticateToken(context *gin.Context) bool {
	token := context.Request.Header.Get("token")

	if token == "" {
		serverResponse.BadRequestServerError(context, "token required")
		context.Abort()
		return false
	}

	jwtData, err := authService.ValidateToken(token)

	if err != "" {
		serverResponse.BadRequestServerError(context, "Invalid token - "+err)
		context.Abort()
		return false
	}
	context.Set("email", jwtData.Email)
	context.Set("name", jwtData.Name)
	context.Set("userName", jwtData.Username)
	context.Set("userId", jwtData.User_Id)
	context.Set("userType", jwtData.User_type)

	return true

}

func Authenticate(context *gin.Context) {
	success := authenticateToken(context)

	if !success {
		return
	}

	context.Next()
}

func AuthenticateAdmin(context *gin.Context) {
	success := authenticateToken(context)

	if !success {
		return
	}

	if context.GetString("userType") != "ADMIN" {
		serverResponse.UnauthorizedRequest(context)
		context.Abort()
		return
	}

	context.Next()
}
