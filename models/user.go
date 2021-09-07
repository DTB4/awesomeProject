package models

import (
	"time"
)

type User struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	SecondName   string    `json:"second_name"`
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
