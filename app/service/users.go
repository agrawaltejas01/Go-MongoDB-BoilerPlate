package service

import (
	"encoding/json"
	"io"

	"github.com/agrawaltejas01/Go-MongoDB-BoilerPlate/app/repositories"
	"github.com/agrawaltejas01/Go-MongoDB-BoilerPlate/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers() ([]db.User, error) {
	return repositories.GetUsers()
}

func GetUser(id string) (db.User, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	return repositories.GetUser(_id)
}

func AddUser(userData io.ReadCloser) (string, error) {
	var user db.User
	json.NewDecoder(userData).Decode(&user)
	objectId, err := repositories.AddUser(user)

	if err != nil {
		return "", err
	}
	return objectId.String(), nil
}

func UpdateUser(id string) error {
	_id, _ := primitive.ObjectIDFromHex(id)
	return repositories.UpdateUser(_id)
}

func DeleteUser(id string) error {
	_id, _ := primitive.ObjectIDFromHex(id)
	return repositories.DeleteUser(_id)
}
