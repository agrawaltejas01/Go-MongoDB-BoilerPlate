package lib

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	//"context"
	//"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func MaskPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func ComparePassword(storedPassword string, passwordEntered string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(passwordEntered))
	check := true

	if err != nil {
		check = false
	}
	return check
}
