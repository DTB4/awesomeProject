package models

import (
	"time"
)

type User struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
	Deleted      bool      `json:"deleted"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ActiveUserData struct {
	ID int `json:"id"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func UserModelTransform(user *User) *UserResponse {
	response := UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return &response
}
