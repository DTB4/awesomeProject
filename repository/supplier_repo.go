package repository

import (
	"awesomeProject/models"
	"database/sql"
	"log"
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
	SearchByID(id int) (bool, error)
	SoftDeleteNotUpdated(interval int) (sql.Result, error)
}

type SupplierRepository struct {
	db *sql.DB
}

func (s SupplierRepository) Create(supplier *models.Supplier) (sql.Result, error) {
	result, err := s.db.Exec("INSERT INTO supliers (id, name, description, created, updated, img_url) VALUES (?, ?, ?, current_timestamp, current_timestamp , ?)", 0, supplier.Name, supplier.Description, supplier.ImgURL)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s SupplierRepository) GetByID(id int) (*models.Supplier, error) {
	supplier := models.Supplier{}
	rows, err := s.db.Query("SELECT * FROM supliers WHERE id=? AND deleted=FALSE", id)
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
	rows, err := s.db.Query("SELECT * FROM supliers WHERE deleted=false")

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
	result, err := s.db.Exec("UPDATE supliers SET name = ?, updated=current_timestamp, img_url=? WHERE id=?", supplier.Name, supplier.ImgURL, supplier.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s SupplierRepository) Delete(id int) (sql.Result, error) {
	result, err := s.db.Exec("DELETE FROM supliers WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s SupplierRepository) SoftDelete(id int) (sql.Result, error) {
	result, err := s.db.Exec("UPDATE supliers SET deleted=true, updated=current_timestamp WHERE id=?", id)
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

func (s SupplierRepository) SoftDeleteALL() (sql.Result, error) {
	result, err := s.db.Exec("UPDATE supliers SET deleted=true, updated=current_timestamp WHERE deleted!=true")
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
func (s SupplierRepository) SoftDeleteNotUpdated(interval int) (sql.Result, error) {
	result, err := s.db.Exec("UPDATE supliers SET deleted=true, updated=current_timestamp WHERE deleted=false AND (current_timestamp-updated )>=?", interval)
	if err != nil {
		return nil, err
	}
	return result, nil
}
