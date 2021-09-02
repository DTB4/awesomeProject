package repository

import (
	"awesomeProject/models"
	"database/sql"
	"log"
	"time"
)

func NewProductsRepository(db *sql.DB) *ProductsRepository {
	return &ProductsRepository{db: db}
}

type ProductsRepositoryI interface {
	CreateNewProduct(product *models.Product) (sql.Result, error)
	GetProductByID(id int) (*models.Product, error)
	GetAllProducts() (*[]models.Product, error)
	GetAllProductsBySupplierID(id int) (*[]models.Product, error)
	EditProduct(product *models.Product) (sql.Result, error)
	DeleteProduct(id int) (sql.Result, error)
	DeleteAllProducts() (sql.Result, error)
	SearchProductBySupIDAndName(supplierID int, name string) (int, error)
}

type ProductsRepository struct {
	db *sql.DB
}

func (p ProductsRepository) CreateNewProduct(product *models.Product) (sql.Result, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	result, err := tx.Exec("INSERT INTO products (name, type, description, price , created, updated, id_supplier, img_url, ingredients) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", product.Name, product.Type, product.Description, product.Price, time.Now(), time.Now(), product.IDSupplier, product.ImgURL, product.Ingredients)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		return nil, err
	}
	return result, nil
}

func (p ProductsRepository) GetProductByID(id int) (*models.Product, error) {
	product := models.Product{}
	rows, err := p.db.Query("SELECT * FROM products WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Type, &product.Description, &product.Price, &product.Created, &product.Updated, &product.Deleted, &product.IDSupplier, &product.ImgURL)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p ProductsRepository) GetAllProducts() (*[]models.Product, error) {
	var products []models.Product
	rows, err := p.db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}
	product := models.Product{}
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name)
		if err != nil {
			log.Println(err)
		}
		products = append(products, product)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (p ProductsRepository) GetAllProductsBySupplierID(id int) (*[]models.Product, error) {
	var products []models.Product
	rows, err := p.db.Query("SELECT * FROM products WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	product := models.Product{}
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Type, &product.Description, &product.Price, &product.Created, &product.Updated, &product.Deleted, &product.IDSupplier, &product.ImgURL, &product.Ingredients)
		if err != nil {
			log.Println(err)
		}
		products = append(products, product)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (p ProductsRepository) EditProduct(product *models.Product) (sql.Result, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	result, err := tx.Exec("UPDATE products SET name=?, type=?, description=?, price=?, updated=?, img_url=?, ingredients=? WHERE id=?", product.Name, product.Type, product.Description, product.Price, time.Now(), product.ImgURL, product.Ingredients, product.ID)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		return nil, err
	}
	return result, nil
}

func (p ProductsRepository) DeleteProduct(id int) (sql.Result, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	result, err := tx.Exec("DELETE from products WHERE id=?", id)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		return nil, err
	}
	return result, nil
}

func (p ProductsRepository) DeleteAllProducts() (sql.Result, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	result, err := tx.Exec("DELETE FROM products")
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		return nil, err
	}
	return result, nil
}

func (p ProductsRepository) SearchProductBySupIDAndName(supplierID int, name string) (int, error) {
	rows, err := p.db.Query("SELECT id FROM products WHERE id_supplier=? AND name=?", supplierID, name)

	if err != nil {
		log.Fatal(err)
	}
	product := models.Product{}
	for rows.Next() {
		err = rows.Scan(&product.ID)
		if err != nil {
			log.Println(err)
		}
		return product.ID, nil
	}
	err = rows.Close()
	if err != nil {
		return 0, err
	}
	return 0, err
}
