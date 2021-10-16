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
	Create(order *models.Order) (int, error)
	GetByID(orderID int) (*models.Order, error)
	GetUserOrders(userID int) (*[]models.Order, error)
	UpdateStatus(order *models.Order) (int, error)
	SetTotal(ID int, total float64) (int, error)
	Delete(id int) (int, error)
}

type OrderRepository struct {
	db *sql.DB
}

func (or OrderRepository) Create(order *models.Order) (int, error) {
	result, err := or.db.Exec("INSERT INTO orders (id, id_user, status, created, updated, adress, contact_number) VALUES (?, ?, ?, ?, ?, ?, ?)", order.ID, order.IDUser, "created", time.Now(), time.Now(), order.Address, order.ContactNumber)
	if err != nil {
		return 0, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastID), nil
}

func (or OrderRepository) GetByID(orderID int) (*models.Order, error) {
	order := models.Order{}
	rows, err := or.db.Query("SELECT * FROM orders WHERE id=? AND deleted=false", orderID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&order.ID, &order.IDUser, &order.Status, &order.Created, &order.Updated, &order.Address, &order.ContactNumber, &order.Total)
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
	rows, err := or.db.Query("SELECT * FROM orders WHERE id_user=? AND deleted=false", userID)
	if err != nil {
		return nil, err
	}
	order := models.Order{}
	for rows.Next() {
		err = rows.Scan(&order.ID, &order.IDUser, &order.Status, &order.Created, &order.Updated, &order.Deleted, &order.Address, &order.ContactNumber, &order.Total)
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

func (or OrderRepository) UpdateStatus(order *models.Order) (int, error) {
	result, err := or.db.Exec("UPDATE orders SET status=?, updated=? WHERE id=? AND deleted=false", order.Status, time.Now(), order.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (or OrderRepository) SetTotal(ID int, total float64) (int, error) {
	result, err := or.db.Exec("UPDATE orders SET total=? WHERE id=? AND deleted=false", total, ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (or OrderRepository) Delete(id int) (int, error) {
	result, err := or.db.Exec("DELETE from orders WHERE id=?", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}
