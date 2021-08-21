package repository

import (
	"awesomeProject/models"
	"database/sql"
)

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

type OrderRepositoryI interface {
	CreateNewOrder(order models.Order) error
	GetOrderByID(orderID int) (models.Order, error)
	GetUserOrders(userID int) ([]models.Order, error)
	GetOrderProducts(orderID int) ([]models.OrderProduct, error)
	CreateOrderProduct(orderID int) (models.OrderProduct, error)
}

type OrderRepository struct {
	db *sql.DB
}

func (o OrderRepository) CreateNewOrder(order models.Order) error {
	panic("implement me")
}

func (o OrderRepository) GetOrderByID(orderID int) (models.Order, error) {
	panic("implement me")
}

func (o OrderRepository) GetUserOrders(userID int) ([]models.Order, error) {
	panic("implement me")
}

func (o OrderRepository) GetOrderProducts(orderID int) ([]models.OrderProduct, error) {
	panic("implement me")
}

func (o OrderRepository) CreateOrderProduct(orderID int) (models.OrderProduct, error) {
	panic("implement me")
}
