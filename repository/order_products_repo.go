package repository

import (
	"awesomeProject/models"
	"database/sql"
	"log"
)

func NewOrderProductsRepository(db *sql.DB) *OrderProductsRepository {
	return &OrderProductsRepository{
		db: db,
	}
}

type OrderProductsRepositoryI interface {
	Create(orderProduct *models.OrderProduct) (sql.Result, error)
	GetByOrderID(id int) (*[]models.OrderProduct, error)
	Update(orderProduct *models.OrderProduct) (sql.Result, error)
	DeleteByIDs(orderID, productID int) (sql.Result, error)
}

type OrderProductsRepository struct {
	db *sql.DB
}

func (op OrderProductsRepository) Create(orderProduct *models.OrderProduct) (sql.Result, error) {
	result, err := op.db.Exec("INSERT INTO order_product (order_id, product_id, quantity) VALUES (?, ?, ?)", orderProduct.OrderID, orderProduct.ProductID, orderProduct.Quantity)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (op OrderProductsRepository) GetByOrderID(id int) (*[]models.OrderProduct, error) {
	var orderProductSlice []models.OrderProduct
	rows, err := op.db.Query("SELECT * FROM order_product WHERE order_id=?", id)
	if err != nil {
		return nil, err
	}
	orderProduct := models.OrderProduct{}
	for rows.Next() {
		err = rows.Scan(&orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity)
		if err != nil {
			log.Println(err)
		}
		orderProductSlice = append(orderProductSlice, orderProduct)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &orderProductSlice, nil
}

func (op OrderProductsRepository) Update(orderProduct *models.OrderProduct) (sql.Result, error) {
	result, err := op.db.Exec("UPDATE order_product SET quantity=? WHERE order_id=? AND product_id=?", orderProduct.Quantity, orderProduct.OrderID, orderProduct.ProductID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (op OrderProductsRepository) DeleteByIDs(orderID, productID int) (sql.Result, error) {
	result, err := op.db.Exec("DELETE from order_product WHERE order_id=? AND product_id=?", orderID, productID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
