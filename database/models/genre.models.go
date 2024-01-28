package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Genre struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name      string             `json:"name" validate:"required,min=3,max=100"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	Genre_id  int                `json:"genre_id" validate:"required"`
}
