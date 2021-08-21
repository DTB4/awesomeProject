package models

import "time"

type Order struct {
	ID      int       `json:"id"`
	IDUser  int       `json:"id_user"`
	Status  string    `json:"status"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
