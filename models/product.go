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
	Deleted     bool      `json:"deleted"`
	IDSupplier  int       `json:"id_supplier"`
	ImgURL      string    `json:"img_url"`
	Ingredients string    `json:"ingredients"`
}

type ResponseBodyMenu struct {
	Menu []ParserProduct `json:"menu"`
}

type ParserProduct struct {
	Image       string   `json:"image"`
	Name        string   `json:"name"`
	Price       float32  `json:"price"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
}

type ProductRequest struct {
	ID int `json:"id"`
}

type ProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Price       float32 `json:"price"`
	ImgURL      string  `json:"img_url"`
	Ingredients string  `json:"ingredients"`
}

type ProductTypeRequest struct {
	ProductType string `json:"product_type"`
}

type ProductSupplierIDRequest struct {
	SupplierID int `json:"supplier_id"`
}
type ProductTypesResponse struct {
	Types []string `json:"types"`
}
