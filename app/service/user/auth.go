package userService

import (
	"errors"
	"fmt"
	userRepo "shive-app/app/repositories/user"
	authService "shive-app/app/service/auth"
	"shive-app/database/models"
	"shive-app/lib"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-playground/validator/v10"
)

func Signup(userObject models.User) (primitive.ObjectID, error) {

	// Check if email and username exists
	emailOrUsernameError := CheckIfEmailOrUsernameIsDuplicate(userObject)
	if emailOrUsernameError != nil {
		return primitive.NilObjectID, emailOrUsernameError
	}

	// Create password hash
	userObject.Password = lib.MaskPassword(userObject.Password)

	// Assign random hex Id as user Id
	userObject.ID = primitive.NewObjectID()
	userObject.User_id = userObject.ID.Hex()

	// Generate token
	token, refreshToken, _ := authService.GenerateAllTokens(
		userObject.Email,
		userObject.Name,
		userObject.Username,
		userObject.User_type,
		userObject.User_id)
	userObject.Token = token
	userObject.Refresh_token = refreshToken

	// Validate struct
	if validationErr := validator.New().Struct(userObject); validationErr != nil {
		return primitive.NilObjectID, errors.New("error in signup validation find operation ")
	}

	// Save
	return userRepo.AddUser(userObject)
}

func Login(email string, password string) (models.User, error) {

	userInDb, findErr := CheckIfEmailExists(email)
	if findErr != nil {
		// For error, userInDb will be {}
		return userInDb, findErr
	}

	// Validate password
	passwordMatches := lib.ComparePassword(userInDb.Password, password)
	if !passwordMatches {
		return models.User{}, errors.New("username or password incorrect")
	}

	// Generate token and refresh token
	token, refreshToken, _ := authService.GenerateAllTokens(
		userInDb.Email,
		userInDb.Name,
		userInDb.Username,
		userInDb.User_type,
		userInDb.User_id)
	// Store refresh token in db
	tokenSaveErr := userRepo.UpdateTokens(userInDb.ID, token, refreshToken)
	if tokenSaveErr != nil {
		fmt.Println("Error in saving tokens")
		return models.User{}, errors.New("error in token")
	}

	// Return user
	return userInDb, nil
}
