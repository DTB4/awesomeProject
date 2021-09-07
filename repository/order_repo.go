package repository

import (
	"awesomeProject/models"
	"database/sql"
	"time"
)

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

type OrderRepositoryI interface {
	Create(order *models.Order) (sql.Result, error)
	GetByID(orderID int) (*models.Order, error)
	GetUserOrders(userID int) (*[]models.Order, error)
	Update(order *models.Order) (sql.Result, error)
	Delete(id int) (sql.Result, error)
}

type OrderRepository struct {
	db *sql.DB
}

func (or OrderRepository) Create(order *models.Order) (sql.Result, error) {
	result, err := or.db.Exec("INSERT INTO orders (id, id_user, status, created) VALUES (?, ?, ?, ?)", order.ID, order.IDUser, "created", time.Now())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (or OrderRepository) GetByID(orderID int) (*models.Order, error) {
	order := models.Order{}
	rows, err := or.db.Query("SELECT * FROM orders WHERE id=?", orderID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&order.ID, &order.IDUser, &order.Status, &order.Created, &order.Updated)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Close()
	if err != nil {
		return nil, err
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
		err = rows.Scan(&order.ID, &order.IDUser, &order.Status, &order.Created, &order.Updated)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &orders, nil
}

func (or OrderRepository) Update(order *models.Order) (sql.Result, error) {
	result, err := or.db.Exec("UPDATE orders SET status=?, updated=? WHERE id=?", order.Status, time.Now(), order.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (or OrderRepository) Delete(id int) (sql.Result, error) {
	result, err := or.db.Exec("DELETE from orders WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
