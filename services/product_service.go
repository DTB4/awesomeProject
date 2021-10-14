package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
)

func NewProductService(productRepository repository.ProductsRepositoryI) *ProductService {
	return &ProductService{
		productRepository,
	}
}

type ProductServiceI interface {
	GetByID(productID int) (*models.Product, error)
	GetAll() (*[]models.Product, error)
	GetAllByType(productType string) (*[]models.Product, error)
	GetAllBySuppliersID(supplierID int) (*[]models.Product, error)
	GetTypes() (*models.ProductTypesResponse, error)
}

type ProductService struct {
	productRepository repository.ProductsRepositoryI
}

func (p ProductService) GetByID(productID int) (*models.Product, error) {
	product, err := p.productRepository.GetByID(productID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p ProductService) GetAll() (*[]models.Product, error) {
	products, err := p.productRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p ProductService) GetAllByType(productType string) (*[]models.Product, error) {
	products, err := p.productRepository.GetALLByType(productType)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p ProductService) GetAllBySuppliersID(supplierID int) (*[]models.Product, error) {
	products, err := p.productRepository.GetAllBySupplierID(supplierID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (p ProductService) GetTypes() (*models.ProductTypesResponse, error) {
	productTypes, err := p.productRepository.GetTypes()
	if err != nil {
		return nil, err
	}
	return productTypes, nil
}
