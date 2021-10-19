package repository

import (
	"awesomeProject/models"
	"database/sql"
	"fmt"
	"log"
)

func NewOrderProductsRepository(db *sql.DB) *OrderProductsRepository {
	return &OrderProductsRepository{
		db: db,
	}
}

type OrderProductsRepositoryI interface {
	Create(orderProduct *models.OrderProduct) (int, error)
	GetByOrderID(id int) (*[]models.OrderProduct, error)
	Update(orderProduct *models.OrderProduct) (int, error)
	DeleteByIDs(orderID, productID int) (int, error)
	GetProductPrice(productID int) (float32, error)
}

type OrderProductsRepository struct {
	db *sql.DB
}

func (op OrderProductsRepository) Create(orderProduct *models.OrderProduct) (int, error) {
	result, err := op.db.Exec("INSERT INTO order_product (id, order_id, product_id, quantity, price, name) VALUES (?, ?, ?, ?, ?, ?)", 0, orderProduct.OrderID, orderProduct.ProductID, orderProduct.Quantity, orderProduct.Price, orderProduct.Name)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (op OrderProductsRepository) GetByOrderID(id int) (*[]models.OrderProduct, error) {
	var orderProductSlice []models.OrderProduct
	rows, err := op.db.Query("SELECT * FROM order_product WHERE order_id=?", id)
	if err != nil {
		return nil, err
	}
	orderProduct := models.OrderProduct{}
	for rows.Next() {
		err = rows.Scan(&orderProduct.ID, &orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity, &orderProduct.Price, &orderProduct.Name)
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

func (op OrderProductsRepository) Update(orderProduct *models.OrderProduct) (int, error) {
	result, err := op.db.Exec("UPDATE order_product SET quantity=?, price=? WHERE id=? ", orderProduct.Quantity, orderProduct.Price, orderProduct.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (op OrderProductsRepository) DeleteByIDs(orderID, productID int) (int, error) {
	result, err := op.db.Exec("DELETE from order_product WHERE order_id=? AND product_id=?", orderID, productID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (op OrderProductsRepository) GetProductPrice(productID int) (float32, error) {
	fmt.Println("repos called")
	var price float32
	rows, err := op.db.Query("SELECT price FROM products WHERE id=?", productID)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		err = rows.Scan(&price)
		if err != nil {
			return 0, err
		}
	}
	fmt.Println("price from repos:", price)
	err = rows.Close()
	if err != nil {
		return price, err
	}
	return price, nil
}
