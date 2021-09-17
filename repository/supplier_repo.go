package repository

import (
	"awesomeProject/models"
	"database/sql"
)

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

type SupplierRepositoryI interface {
	Create(restaurant *models.Supplier) (int, error)
	GetByID(id int) (*models.Supplier, error)
	GetByName(name string) (*models.Supplier, error)
	GetAll() (*[]models.Supplier, error)
	GetAllByType(supplierType string) (*[]models.Supplier, error)
	GetAllByTime(time string) (*[]models.Supplier, error)
	Update(restaurant *models.Supplier) (int, error)
	Delete(id int) (int, error)
	SoftDelete(id int) (int, error)
	Truncate() (int, error)
	SearchByID(id int) (bool, error)
	SoftDeleteNotUpdated(interval int) (int, error)
}

type SupplierRepository struct {
	db *sql.DB
}

func (s SupplierRepository) Create(supplier *models.Supplier) (int, error) {
	result, err := s.db.Exec("INSERT INTO suppliers (id, name, created, updated, img_url, type, opening, closing) VALUES (?, ?, current_timestamp, current_timestamp , ?, ?, ?, ?)", 0, supplier.Name, supplier.ImgURL, supplier.Type, supplier.Opening, supplier.Closing)
	if err != nil {
		return 0, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastID), nil
}

func (s SupplierRepository) GetByID(id int) (*models.Supplier, error) {
	supplier := models.Supplier{}
	rows, err := s.db.Query("SELECT * FROM suppliers WHERE id=? AND deleted=FALSE", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Created, &supplier.Updated, &supplier.Deleted, &supplier.ImgURL, &supplier.Type, &supplier.Opening, &supplier.Closing)
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
	rows, err := s.db.Query("SELECT * FROM suppliers WHERE name=?", name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Created, &supplier.Updated, &supplier.Deleted, &supplier.ImgURL, &supplier.Type, &supplier.Opening, &supplier.Closing)
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
	rows, err := s.db.Query("SELECT * FROM suppliers WHERE deleted=false")

	if err != nil {
		return nil, err
	}
	supplier := models.Supplier{}
	for rows.Next() {
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Created, &supplier.Updated, &supplier.Deleted, &supplier.ImgURL, &supplier.Type, &supplier.Opening, &supplier.Closing)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &suppliers, nil
}

func (s SupplierRepository) GetAllByType(supplierType string) (*[]models.Supplier, error) {
	var suppliers []models.Supplier
	rows, err := s.db.Query("SELECT * FROM suppliers WHERE deleted=false AND type=?", supplierType)

	if err != nil {
		return nil, err
	}
	supplier := models.Supplier{}
	for rows.Next() {
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Created, &supplier.Updated, &supplier.Deleted, &supplier.ImgURL, &supplier.Type, &supplier.Opening, &supplier.Closing)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &suppliers, nil
}
func (s SupplierRepository) GetAllByTime(time string) (*[]models.Supplier, error) {
	var suppliers []models.Supplier
	rows, err := s.db.Query("SELECT * FROM suppliers WHERE deleted=false AND opening<=? AND closing>?", time, time)

	if err != nil {
		return nil, err
	}
	supplier := models.Supplier{}
	for rows.Next() {
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Created, &supplier.Updated, &supplier.Deleted, &supplier.ImgURL, &supplier.Type, &supplier.Opening, &supplier.Closing)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &suppliers, nil
}

func (s SupplierRepository) Update(supplier *models.Supplier) (int, error) {
	result, err := s.db.Exec("UPDATE suppliers SET name = ?, updated=current_timestamp, deleted=false, img_url=?, type=?, opening=?, closing=? WHERE id=?", supplier.Name, supplier.ImgURL, supplier.Type, supplier.Opening, supplier.Closing, supplier.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (s SupplierRepository) Delete(id int) (int, error) {
	result, err := s.db.Exec("DELETE FROM suppliers WHERE id=?", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (s SupplierRepository) SoftDelete(id int) (int, error) {
	result, err := s.db.Exec("UPDATE suppliers SET deleted=true, updated=current_timestamp WHERE id=?", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (s SupplierRepository) Truncate() (int, error) {
	result, err := s.db.Exec("DELETE FROM suppliers")
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (s SupplierRepository) SoftDeleteALL() (int, error) {
	result, err := s.db.Exec("UPDATE suppliers SET deleted=true, updated=current_timestamp WHERE deleted!=true")
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (s SupplierRepository) SearchByID(id int) (bool, error) {
	rows, err := s.db.Query("SELECT * FROM suppliers WHERE id=?", id)
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
func (s SupplierRepository) SoftDeleteNotUpdated(interval int) (int, error) {
	result, err := s.db.Exec("UPDATE suppliers SET deleted=true, updated=current_timestamp WHERE deleted=false AND (current_timestamp-updated )>=?", interval)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}
