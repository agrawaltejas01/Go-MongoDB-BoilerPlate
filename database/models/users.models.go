package models

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id, omitempty"`
	Name          string             `json:"name" validate:"required,min=4,max=100"`
	Username      string             `json:"username" validate:"required,min=4,max=100"`
	Password      string             `json:"password" validate:"required,min=8"`
	Email         string             `json:"email" validate:"email,required"`
	Token         string             `json:"token"`
	User_type     string             `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token string             `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id" validate: "required, min=4"`
}

func marshal(u *User) *User {

	if u.ID.IsZero() {
		u.ID = primitive.NewObjectID()
	}
	if u.Created_at.IsZero() {
		u.Created_at = time.Now()
	}
	u.Updated_at = time.Now()

	return u

	// type my User
	// return bson.Marshal((*my)(u))
}

func (u *User) MarshalBSON() ([]byte, error) {

	u = marshal(u)

	type my User
	return bson.Marshal((*my)(u))
}

func (u *User) MarshalJSON() ([]byte, error) {
	u = marshal(u)

	type my User
	return json.Marshal((*my)(u))
}
