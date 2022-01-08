package service

import (
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
