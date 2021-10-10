package models

import "time"

type Order struct {
	ID      int       `json:"id"`
	IDUser  int       `json:"id_user"`
	Status  string    `json:"status"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Deleted bool      `json:"deleted"`
}

type OrderProduct struct {
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type RequestOrderID struct {
	OrderId int `json:"order_id"`
}

type UpdateOrderRequest struct {
	OrderID int    `json:"order_id"`
	Status  string `json:"status"`
}

type OrderCreationResponse struct {
	OrderID    int `json:"order_id"`
	ProductQty int `json:"product_qty"`
}
