package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type user struct {
	ID            primitive.ObjectID `bson:"id"`
	First_Name    *string            `json:"first_name" validate:"required" min:"2" max:"100"`
	Last_Name     *string            `json:"last_name" validate:"required" min:"2" max:"100"`
	Phone         *string            `json:"phone" validate:"required" min:"10"`
	Email         *string            `json:"email" validate:"email, required"`
	Password      *string            `json:"password" validate:"required, min = 8, max = 10, uppercase, lowercase, specialchar, number"`
	User          *string            `json:"user"`
	User_Type     *string            `json:"user_type" validate:"required, eq=USER|eq=ADMIN"`
	Token         *string            `json:"token"`
	Refresh_Token *string            `json:"refresh_token"`
	Created_At    time.Time          `json:"created_at"`
	Updated_At    time.Time          `json:"updated_at"`
}
