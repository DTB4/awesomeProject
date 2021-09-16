package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
)

func NewSupplierService(supplierRepository repository.SupplierRepositoryI) *SupplierService {
	return &SupplierService{
		supplierRepository,
	}
}

type SupplierServiceI interface {
	GetByID(supID int) (*models.Supplier, error)
	GetAll() (*[]models.Supplier, error)
}

type SupplierService struct {
	supplierRepository repository.SupplierRepositoryI
}

func (s SupplierService) GetByID(supID int) (*models.Supplier, error) {
	supplier, err := s.supplierRepository.GetByID(supID)
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s SupplierService) GetAll() (*[]models.Supplier, error) {
	suppliers, err := s.supplierRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}
