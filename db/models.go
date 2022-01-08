package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
}
