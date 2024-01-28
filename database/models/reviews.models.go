package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	Id         primitive.ObjectID `bson:"_id"`
	Movie_id   int                `json:"movie_id" validate:"required"`
	Review     string             `json:"review" validate:"required"`
	Review_id  string             `json:"review_id" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	User_id    string             `json:"user_id" validate:"required"`
}
