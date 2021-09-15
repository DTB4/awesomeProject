package models

import (
	"errors"
	"strings"
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

//UserValidate for fast and simple validation of name and email
func UserValidate(user User) (bool, error) {
	if len(user.FirstName) < 3 {
		return false, errors.New("wrong firstname")
	}
	if len(user.LastName) < 3 {
		return false, errors.New("wrong lastname")
	}
	if len(user.Email) < 7 || !strings.Contains(user.Email, "@") {
		return false, errors.New("invalid email")
	}
	return true, nil
}
