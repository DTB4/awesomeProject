package models

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
