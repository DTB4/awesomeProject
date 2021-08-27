package repository

import (
	"awesomeProject/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

type OrderRepositoryI interface {
	CreateNewOrder(order *models.Order) error
	FillOrderWithProducts(OrderID int, products *[]models.Product) error
	GetOrderByID(orderID int) (*models.Order, error)
	GetUserOrders(userID int) (*[]models.Order, error)
	GetOrderProducts(orderID int) (*[]models.OrderProduct, error)
	CreateOrderProducts(orderProducts *[]models.OrderProduct) error
}

type OrderRepository struct {
	db *sql.DB
}

func (or OrderRepository) CreateNewOrder(order *models.Order) error {
	result, err := or.db.Exec("INSERT INTO orders (id, id_user, status, created) VALUES (?, ?, ?, ?)", order.ID, order.IDUser, "created", time.Now())
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

func (or OrderRepository) GetOrderByID(orderID int) (*models.Order, error) {
	order := models.Order{}
	rows, err := or.db.Query("SELECT * FROM orders WHERE id=?", orderID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&order.ID, &order.IDUser, &order.Status, &order.Created, &order.Updated)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(order)
	}
	return &order, nil
}

func (or OrderRepository) GetUserOrders(userID int) (*[]models.Order, error) {
	var orders []models.Order
	rows, err := or.db.Query("SELECT * FROM orders WHERE id_user=?", userID)
	if err != nil {
		return nil, err
	}
	order := models.Order{}
	for rows.Next() {
		err := rows.Scan(&order.ID, &order.IDUser, &order.Status, &order.Created, &order.Updated)
		if err != nil {
			log.Println(err)
		}
		orders = append(orders, order)
	}
	return &orders, nil
}

func (or OrderRepository) GetOrderProducts(orderID int) (*[]models.OrderProduct, error) {
	var orderProductSlice []models.OrderProduct
	rows, err := or.db.Query("SELECT * FROM order_product WHERE order_id=?", orderID)
	if err != nil {
		return nil, err
	}
	orderProduct := models.OrderProduct{}
	for rows.Next() {
		err := rows.Scan(&orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity)
		if err != nil {
			log.Println(err)
		}
		orderProductSlice = append(orderProductSlice, orderProduct)
	}
	return &orderProductSlice, nil
}

func (or OrderRepository) CreateOrderProducts(orderProducts *[]models.OrderProduct) error {
	transaction, err := or.db.Begin()
	if err != nil {
		return err
	}
	for i := 0; i <= len(*orderProducts); i++ {
		result, err := transaction.Exec("INSERT INTO order_product (order_id, product_id, quantity) VALUES (?, ?, ?)", (*orderProducts)[i].OrderID, (*orderProducts)[i].ProductID, (*orderProducts)[i].Quantity)
		if err != nil {
			err := transaction.Rollback()
			return err
		}
		log.Println(result)
	}
	err = transaction.Commit()
	err = or.db.Close()
	if err != nil {
		return err
	}
	return nil
}
