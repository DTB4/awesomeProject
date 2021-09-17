package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
	"errors"
)

func NewSupplierService(supplierRepository repository.SupplierRepositoryI) *SupplierService {
	return &SupplierService{
		supplierRepository,
	}
}

type SupplierServiceI interface {
	GetByID(supID int) (*models.SupplierResponse, error)
	GetAll() (*[]models.SupplierResponse, error)
	GetAllByType(supplierType string) (*[]models.SupplierResponse, error)
	GetAllByTime(time string) (*[]models.SupplierResponse, error)
}

type SupplierService struct {
	supplierRepository repository.SupplierRepositoryI
}

func (s SupplierService) GetByID(supID int) (*models.SupplierResponse, error) {
	supplier, err := s.supplierRepository.GetByID(supID)
	if err != nil {
		return nil, err
	}
	supplierResponse := models.TransformSupplierForResponse(supplier)
	return supplierResponse, nil
}

func (s SupplierService) GetAll() (*[]models.SupplierResponse, error) {
	var suppliersResponse []models.SupplierResponse
	suppliers, err := s.supplierRepository.GetAll()
	if err != nil {
		return nil, err
	}
	for _, supplier := range *suppliers {
		s := models.TransformSupplierForResponse(&supplier)
		suppliersResponse = append(suppliersResponse, *s)
	}
	return &suppliersResponse, nil
}

func (s SupplierService) GetAllByType(supplierType string) (*[]models.SupplierResponse, error) {
	var suppliersResponse []models.SupplierResponse
	suppliers, err := s.supplierRepository.GetAllByType(supplierType)
	if err != nil {
		return nil, err
	}
	if suppliers == nil {
		return nil, errors.New("no suppliers of this type")
	}
	if err != nil {
		return nil, err
	}
	for _, supplier := range *suppliers {
		s := models.TransformSupplierForResponse(&supplier)
		suppliersResponse = append(suppliersResponse, *s)
	}
	return &suppliersResponse, nil
}
func (s SupplierService) GetAllByTime(time string) (*[]models.SupplierResponse, error) {
	var suppliersResponse []models.SupplierResponse
	suppliers, err := s.supplierRepository.GetAllByTime(time)
	if err != nil {
		return nil, err
	}
	if suppliers == nil {
		return nil, errors.New("no suppliers open in this time")
	}
	for _, supplier := range *suppliers {
		s := models.TransformSupplierForResponse(&supplier)
		suppliersResponse = append(suppliersResponse, *s)
	}
	return &suppliersResponse, nil
}
