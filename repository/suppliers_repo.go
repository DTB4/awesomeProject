package repository

import (
	"awesomeProject/models"
	"database/sql"
	"log"
	"time"
)

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

type SupplierRepositoryI interface {
	CreateNewSupplier(restaurant *models.Supplier) (sql.Result, error)
	GetSupplierByID(id int) (*models.Supplier, error)
	GetAllSuppliers() (*[]models.Supplier, error)
	EditSupplier(restaurant *models.Supplier) (sql.Result, error)
	DeleteSupplier(id int) (sql.Result, error)
	DeleteAllSuppliers() (sql.Result, error)
	SearchSupplierByID(id int) (bool, error)
}

type SupplierRepository struct {
	db *sql.DB
}

func (s SupplierRepository) CreateNewSupplier(supplier *models.Supplier) (sql.Result, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	result, err := tx.Exec("INSERT INTO supliers (id, name, description, created, updated, img_url) VALUES (?, ?, ?, ?, ?, ?)", supplier.ID, supplier.Name, supplier.Description, time.Now(), time.Now(), supplier.ImgURL)
	if err != nil {
		err = tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		return nil, err
	}
	return result, nil
}

func (s SupplierRepository) GetSupplierByID(id int) (*models.Supplier, error) {
	supplier := models.Supplier{}
	rows, err := s.db.Query("SELECT * FROM supliers WHERE id=?", id)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Description, &supplier.Created, &supplier.Updated, &supplier.Deleted, supplier.ImgURL)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (s SupplierRepository) GetAllSuppliers() (*[]models.Supplier, error) {
	var suppliers []models.Supplier
	rows, err := s.db.Query("SELECT * FROM supliers")

	if err != nil {
		log.Fatal(err)
	}
	supplier := models.Supplier{}
	for rows.Next() {
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Description, &supplier.Created, &supplier.Updated, &supplier.Deleted, supplier.ImgURL)
		if err != nil {
			log.Println(err)
		}
		suppliers = append(suppliers, supplier)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &suppliers, nil
}

func (s SupplierRepository) EditSupplier(supplier *models.Supplier) (sql.Result, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	result, err := tx.Exec("UPDATE supliers SET name = ?, updated=? WHERE id=?", supplier.Name, time.Now(), supplier.ID)
	if err != nil {
		err = tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		return nil, err
	}
	return result, nil
}

func (s SupplierRepository) DeleteSupplier(id int) (sql.Result, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	result, err := tx.Exec("DELETE FROM supliers WHERE id=?", id)
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

func (s SupplierRepository) DeleteAllSuppliers() (sql.Result, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	result, err := tx.Exec("DELETE FROM supliers")
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

func (s SupplierRepository) SearchSupplierByID(id int) (bool, error) {
	rows, err := s.db.Query("SELECT * FROM supliers WHERE id=?", id)

	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	err = rows.Close()
	if err != nil {
		return false, err
	}
	return false, nil
}
