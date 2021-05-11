package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name" validate:"required"`
	Lastname string             `json:"lastname" validate:"required"`
	Age      int32              `json:"age" validate:"required"`
	Email    string             `json:"email" validate:"email"`
	Gender   string             `json:"gender,omitempty"`
}
