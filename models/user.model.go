package models

import (
	"github.com/golang-jwt/jwt"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email       string `json:"email"`
	AccountType string `json:"accountType"`
	jwt.StandardClaims
}

type Session struct {
	Email       string `json:"email"`
	Token       string `json:"token"`
	AccountType string `json:"accountType"`
}

type UserType string

const (
	AdminUser   UserType = "admin"
	DefaultUser UserType = "default"
)

// @MongoDB User
type User struct {
	ID          *string  `json:"id,omitempty" bson:"_id,omitempty" validate:"omitempty"`
	Email       string   `json:"email" bson:"email" validate:"required,email" unique:"true"`
	Password    string   `json:"-" bson:"password" validate:"required" omitempty:"true"`
	AccountType UserType `json:"accountType" bson:"accountType" validate:"required,oneof=admin default"`
	FirstName   *string  `json:"firstName" bson:"firstName" omitempty:"true"`
	LastName    *string  `json:"lastName" bson:"lastName" omitempty:"true"`
	Phone       *string  `json:"phone" bson:"phone" omitempty:"true"`
	Address     *string  `json:"address" bson:"address" omitempty:"true"`
	City        *string  `json:"city" bson:"city" omitempty:"true"`
	Country     *string  `json:"country" bson:"country" omitempty:"true"`
	PostalCode  *string  `json:"postalCode" bson:"postalCode" omitempty:"true"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email     string  `json:"email" validate:"required,email" unique:"true"`
	Password  string  `json:"password" validate:"required"`
	FirstName *string `json:"firstName" omitempty:"true"`
	LastName  *string `json:"lastName" omitempty:"true"`
}
