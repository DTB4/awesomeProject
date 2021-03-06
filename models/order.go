package models

import "time"

type Order struct {
	ID            int       `json:"id"`
	Address       string    `json:"address"`
	ContactNumber string    `json:"contact_number"`
	IDUser        int       `json:"id_user"`
	Status        string    `json:"status"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
	Deleted       bool      `json:"deleted"`
	Total         float64   `json:"total"`
}

type OrderRequest struct {
	Address       string         `json:"address"`
	ContactNumber string         `json:"contact_number"`
	Products      []OrderProduct `json:"products"`
}

type OrderProduct struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
	Name      string  `json:"name"`
}

type RequestOrderID struct {
	OrderId int `json:"order_id"`
}

type UpdateOrderRequest struct {
	OrderID int    `json:"order_id"`
	Status  string `json:"status"`
}

type OrderCreationResponse struct {
	OrderID    int     `json:"order_id"`
	ProductQty int     `json:"product_qty"`
	Total      float64 `json:"total"`
}
