package controllers

import (
	userService "shive-app/app/service/user"
	"shive-app/database/models"
	serverResponse "shive-app/lib/server-response"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {

	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var request requestBody
	jsonBindErr := context.BindJSON(&request)

	if jsonBindErr != nil {
		serverResponse.BadRequestServerError(context, "Bad Data received for user object "+jsonBindErr.Error())
		return
	}

	if request.Email == "" || request.Password == "" {
		serverResponse.BadRequestServerError(context, "Email and Password are required")
		return
	}

	userObject, loginErr := userService.Login(request.Email, request.Password)

	if loginErr != nil {
		serverResponse.InternalServerError(context, "Error in Login - "+loginErr.Error())
		return
	}

	var userData = make(map[string]interface{})
	userData["email"] = userObject.Email
	userData["name"] = userObject.Name
	userData["username"] = userObject.Username
	userData["token"] = userObject.Token
	userData["refresh_token"] = userObject.Refresh_token
	userData["user_id"] = userObject.User_id
	userData["user_type"] = userObject.User_type

	var result = make(map[string]interface{})
	result["user"] = userData
	serverResponse.SuccessResponse(context, result, 0)
}

func Signup(context *gin.Context) {

	var user models.User
	jsonBindErr := context.BindJSON(&user)

	if jsonBindErr != nil {
		serverResponse.BadRequestServerError(context, "Bad Data received for user object "+jsonBindErr.Error())
		return
	}

	id, creationErr := userService.Signup(user)

	if creationErr != nil {
		serverResponse.InternalServerError(context, "Error in Creating User - "+creationErr.Error())
		return
	}

	var result = make(map[string]interface{})
	result["id"] = id

	serverResponse.SuccessResponse(context, result, 0)
}
