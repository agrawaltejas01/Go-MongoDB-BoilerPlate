package userService

import (
	"errors"
	userRepo "shive-app/app/repositories/user"
	"shive-app/database/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func CheckIfEmailOrUsernameAlreadyExists(userObject models.User) error {
	emailOrUsernameUser, emailOrUsernameUserFindErr := userRepo.FindByEmailOrUserName(userObject.Email, userObject.Username)

	if emailOrUsernameUserFindErr != nil && emailOrUsernameUserFindErr != mongo.ErrNoDocuments {
		return errors.New("error in signup validation find operation ")
	}

	if emailOrUsernameUser != (models.User{}) {
		if userObject.Email == emailOrUsernameUser.Email {
			return errors.New("email already present in db")
		} else {
			return errors.New("username already present in db")
		}
	}

	return nil
}
