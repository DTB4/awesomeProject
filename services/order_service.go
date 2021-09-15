package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
	"database/sql"
)

func NewOrderService(orderRepository repository.OrderRepositoryI, orderProductsRepository repository.OrderProductsRepositoryI) *OrderService {
	return &OrderService{
		orderRepository,
		orderProductsRepository,
	}
}

type OrderServiceI interface {
	CreateOrder(order *models.Order) (sql.Result, error)
	CreateOrderProducts(orderID int, orderProducts *[]models.OrderProduct) (int, error)
	GetByID(orderID int) (*[]models.OrderProduct, error)
	GetAll(userID int) (*[]models.Order, error)
	Update(updateRequest *models.UpdateOrderRequest) (sql.Result, error)
	Delete(orderID int) (sql.Result, error)
}

type OrderService struct {
	orderRepository         repository.OrderRepositoryI
	orderProductsRepository repository.OrderProductsRepositoryI
}

func (o OrderService) CreateOrder(order *models.Order) (sql.Result, error) {
	result, err := o.orderRepository.Create(order)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o OrderService) CreateOrderProducts(orderID int, orderProducts *[]models.OrderProduct) (int, error) {
	var iResult int64
	for i := range *orderProducts {
		(*orderProducts)[i].OrderID = orderID
		result, err := o.orderProductsRepository.Create(&(*orderProducts)[i])
		if err != nil {
			return 0, err
		}
		rowsAffected, err := result.RowsAffected()
		iResult += rowsAffected
	}
	return int(iResult), nil
}

func (o OrderService) GetByID(orderID int) (*[]models.OrderProduct, error) {
	orderProducts, err := o.orderProductsRepository.GetByOrderID(orderID)
	if err != nil {
		return nil, err
	}
	return orderProducts, nil
}

func (o OrderService) GetAll(userID int) (*[]models.Order, error) {
	orders, err := o.orderRepository.GetUserOrders(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o OrderService) Update(updateRequest *models.UpdateOrderRequest) (sql.Result, error) {
	var order = new(models.Order)
	order.ID = updateRequest.OrderID
	if updateRequest.Status != "" {
		order.Status = updateRequest.Status
	}
	result, err := o.orderRepository.Update(order)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o OrderService) Delete(orderID int) (sql.Result, error) {
	result, err := o.orderRepository.Delete(orderID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
