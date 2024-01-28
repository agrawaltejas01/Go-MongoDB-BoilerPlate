package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Genre struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      *string            `json:"name" validate:"required,min=4,max=100"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	Genre_id  int                `json:"genreId" validate:"required"`
}
