package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type SignedDetails struct {
	Email string `json:"email,omitempty" validate:"required"`
	Name  string `json:"name,omitempty" validate:"required"`
	Id    string `json:"id,omitempty"`
	jwt.StandardClaims
}
