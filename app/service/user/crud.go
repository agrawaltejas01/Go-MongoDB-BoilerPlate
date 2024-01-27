package userService

import (
	"errors"
	userRepo "shive-app/app/repositories/user"
	"shive-app/database/models"
	"shive-app/lib"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-playground/validator/v10"
)

func Create(userObject models.User) (primitive.ObjectID, error) {

	// Check if email and username exists
	emailOrUsernameError := CheckIfEmailOrUsernameAlreadyExists(userObject)
	if emailOrUsernameError != nil {
		return primitive.NilObjectID, emailOrUsernameError
	}

	// Create password hash
	userObject.Password = lib.MaskPassword(userObject.Password)

	// Assign random hex Id as user Id
	userObject.ID = primitive.NewObjectID()
	// userObject.User_id = userObject.ID.Hex()

	// Generate token
	token, refreshToken, _ := lib.GenerateAllTokens(
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
