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
	Create(restaurant *models.Supplier) (sql.Result, error)
	GetByID(id int) (*models.Supplier, error)
	GetByName(name string) (*models.Supplier, error)
	GetAll() (*[]models.Supplier, error)
	Update(restaurant *models.Supplier) (sql.Result, error)
	Delete(id int) (sql.Result, error)
	SoftDelete(id int) (sql.Result, error)
	Truncate() (sql.Result, error)
	SoftDeleteAll() (sql.Result, error)
	SearchByID(id int) (bool, error)
}

type SupplierRepository struct {
	db *sql.DB
}

func (s SupplierRepository) Create(supplier *models.Supplier) (sql.Result, error) {
	result, err := s.db.Exec("INSERT INTO supliers (name, description, created, updated, img_url) VALUES (?, ?, ?, ?, ?)", supplier.Name, supplier.Description, time.Now(), time.Now(), supplier.ImgURL)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s SupplierRepository) GetByID(id int) (*models.Supplier, error) {
	supplier := models.Supplier{}
	rows, err := s.db.Query("SELECT * FROM supliers WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Description, &supplier.Created, &supplier.Updated, &supplier.Deleted, &supplier.ImgURL)
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

func (s SupplierRepository) GetByName(name string) (*models.Supplier, error) {
	supplier := models.Supplier{}
	rows, err := s.db.Query("SELECT * FROM supliers WHERE name=?", name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Description, &supplier.Created, &supplier.Updated, &supplier.Deleted, &supplier.ImgURL)
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

func (s SupplierRepository) GetAll() (*[]models.Supplier, error) {
	var suppliers []models.Supplier
	rows, err := s.db.Query("SELECT * FROM supliers")

	if err != nil {
		log.Fatal(err)
	}
	supplier := models.Supplier{}
	for rows.Next() {
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Description, &supplier.Created, &supplier.Updated, &supplier.Deleted, &supplier.ImgURL)
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

func (s SupplierRepository) Update(supplier *models.Supplier) (sql.Result, error) {
	result, err := s.db.Exec("UPDATE supliers SET name = ?, updated=?, img_url=? WHERE id=?", supplier.Name, time.Now(), supplier.ImgURL, supplier.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s SupplierRepository) Delete(id int) (sql.Result, error) {
	result, err := s.db.Exec("DELETE from supliers WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s SupplierRepository) SoftDelete(id int) (sql.Result, error) {
	result, err := s.db.Exec("UPDATE supliers SET deleted=true, updated=current_time WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s SupplierRepository) Truncate() (sql.Result, error) {
	result, err := s.db.Exec("DELETE FROM supliers")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s SupplierRepository) SoftDeleteAll() (sql.Result, error) {
	result, err := s.db.Exec("UPDATE supliers SET deleted=true, updated=current_time WHERE deleted!=true")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s SupplierRepository) SearchByID(id int) (bool, error) {
	rows, err := s.db.Query("SELECT * FROM supliers WHERE id=?", id)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		err = rows.Close()
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, nil
}

func (s SupplierRepository) SearchByName(name string) (bool, error) {
	rows, err := s.db.Query("SELECT * FROM supliers WHERE name=?", name)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		err = rows.Close()
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, nil
}
