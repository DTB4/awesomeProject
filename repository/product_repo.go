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
	Create(product *models.Product) (sql.Result, error)
	GetByID(id int) (*models.Product, error)
	GetAll() (*[]models.Product, error)
	GetAllBySupplierID(id int) (*[]models.Product, error)
	GetALLByType(typ string) (*[]models.Product, error)
	Update(product *models.Product) (sql.Result, error)
	Delete(id int) (sql.Result, error)
	SoftDelete(id int) (sql.Result, error)
	Truncate() (sql.Result, error)
	SoftDeleteALL() (sql.Result, error)
	SearchBySupIDAndName(supplierID int, name string) (int, error)
	SoftDeleteNotUpdated(interval int) (sql.Result, error)
}

type ProductsRepository struct {
	db *sql.DB
}

func (p ProductsRepository) Create(product *models.Product) (sql.Result, error) {
	result, err := p.db.Exec("INSERT INTO products (id, name, type, description, price , created, updated, id_supplier, img_url, ingredients) VALUES (?, ?, ?, ?, ?, current_timestamp, current_timestamp, ?, ?, ?)", 0, product.Name, product.Type, product.Description, product.Price, product.IDSupplier, product.ImgURL, product.Ingredients)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p ProductsRepository) GetByID(id int) (*models.Product, error) {
	product := models.Product{}
	rows, err := p.db.Query("SELECT * FROM products WHERE id=? AND deleted=false", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Type, &product.Description, &product.Price, &product.Created, &product.Updated, &product.Deleted, &product.IDSupplier, &product.ImgURL, &product.Ingredients)
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

func (p ProductsRepository) GetAll() (*[]models.Product, error) {
	var products []models.Product
	rows, err := p.db.Query("SELECT * FROM products WHERE deleted=false")
	if err != nil {
		return nil, err
	}
	pr := models.Product{}
	for rows.Next() {
		err = rows.Scan(&pr.ID, &pr.Name, &pr.Type, &pr.Description, &pr.Price, &pr.Created, &pr.Updated, &pr.Deleted, &pr.IDSupplier, &pr.ImgURL, &pr.Ingredients)
		if err != nil {
			log.Println(err)
		}
		products = append(products, pr)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (p ProductsRepository) GetAllBySupplierID(id int) (*[]models.Product, error) {
	var products []models.Product
	rows, err := p.db.Query("SELECT * FROM products WHERE id_supplier=? AND deleted=false", id)
	if err != nil {
		return nil, err
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

func (p ProductsRepository) GetALLByType(productType string) (*[]models.Product, error) {
	var products []models.Product
	rows, err := p.db.Query("SELECT * FROM products WHERE type=? AND deleted=false", productType)
	if err != nil {
		return nil, err
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

func (p ProductsRepository) Update(product *models.Product) (sql.Result, error) {
	result, err := p.db.Exec("UPDATE products SET name=?, type=?, description=?, price=?, updated=current_timestamp, img_url=?, ingredients=? WHERE id=?", product.Name, product.Type, product.Description, product.Price, time.Now(), product.ImgURL, product.Ingredients, product.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p ProductsRepository) SoftDelete(id int) (sql.Result, error) {
	result, err := p.db.Exec("UPDATE products SET deleted=true, updated=current_timestamp WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p ProductsRepository) Delete(id int) (sql.Result, error) {
	result, err := p.db.Exec("DELETE FROM products WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p ProductsRepository) Truncate() (sql.Result, error) {
	result, err := p.db.Exec("DELETE FROM products")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p ProductsRepository) SoftDeleteALL() (sql.Result, error) {
	result, err := p.db.Exec("UPDATE products SET deleted=true, updated=current_timestamp WHERE deleted!=true")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p ProductsRepository) SearchBySupIDAndName(supplierID int, name string) (int, error) {
	rows, err := p.db.Query("SELECT id FROM products WHERE id_supplier=? AND name=?", supplierID, name)
	if err != nil {
		return 0, err
	}
	product := models.Product{}
	for rows.Next() {
		err = rows.Scan(&product.ID)
		if err != nil {
			log.Println(err)
		}
	}
	err = rows.Close()
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (p ProductsRepository) SoftDeleteNotUpdated(interval int) (sql.Result, error) {
	result, err := p.db.Exec("UPDATE products SET deleted=true, updated=current_timestamp WHERE deleted=false AND (current_timestamp-updated )>=?", interval)
	if err != nil {
		return nil, err
	}
	return result, nil
}
