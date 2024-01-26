package userService

import (
	"fmt"
	userRepo "shive-app/app/repositories/user"
	"shive-app/database/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create() (primitive.ObjectID, error) {

	userName := "Tejas"
	email := "tejas@creditail.in"
	userType := "ADMIN"

	userObject := models.User{
		Name:      &userName,
		Username:  &userName,
		Password:  &userName,
		Email:     &email,
		User_type: &userType,
	}

	fmt.Println("Hello User Service")

	return userRepo.AddUser(userObject)

	// return errors.New("test error")
}
