package models

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	Deleted     time.Time `json:"deleted"`
	IDSupplier  int       `json:"id_supplier"`
	ImgURL      string    `json:"img_url"`
	Ingredients string    `json:"ingredients"`
}

type ResponseBodyMenu struct {
	Menu []ParserProduct `json:"menu"`
}

type ParserProduct struct {
	Name        string   `json:"name"`
	Price       float32  `json:"price"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
}
