package models

import "time"

type Supplier struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	Deleted     time.Time `json:"deleted"`
	ImgURL      string    `json:"img_url"`
}

type ResponseBodyRestaurants struct {
	Restaurants []ParserRestaurant `json:"restaurants"`
}

type ParserRestaurant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
