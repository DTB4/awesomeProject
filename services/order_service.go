package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
)

func NewOrderService(orderRepository repository.OrderRepositoryI, orderProductsRepository repository.OrderProductsRepositoryI) *OrderService {
	return &OrderService{
		orderRepository,
		orderProductsRepository,
	}
}

type OrderServiceI interface {
	CreateOrder(order *models.Order) (int, error)
	CreateOrderProducts(orderID int, orderProducts *[]models.OrderProduct) (int, error)
	GetByID(orderID int) (*[]models.OrderProduct, error)
	GetAll(userID int) (*[]models.Order, error)
	Update(updateRequest *models.UpdateOrderRequest) (int, error)
	Delete(orderID int) (int, error)
}

type OrderService struct {
	orderRepository         repository.OrderRepositoryI
	orderProductsRepository repository.OrderProductsRepositoryI
}

func (o OrderService) CreateOrder(order *models.Order) (int, error) {
	orderID, err := o.orderRepository.Create(order)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func (o OrderService) CreateOrderProducts(orderID int, orderProducts *[]models.OrderProduct) (int, error) {
	var rowsAffectedTotal int
	for i := range *orderProducts {
		(*orderProducts)[i].OrderID = orderID
		rowsAffected, err := o.orderProductsRepository.Create(&(*orderProducts)[i])
		if err != nil {
			return 0, err
		}
		rowsAffectedTotal += rowsAffected
	}
	return rowsAffectedTotal, nil
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

func (o OrderService) Update(updateRequest *models.UpdateOrderRequest) (int, error) {
	var order = new(models.Order)
	order.ID = updateRequest.OrderID
	if updateRequest.Status != "" {
		order.Status = updateRequest.Status
	}
	rowsAffected, err := o.orderRepository.Update(order)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (o OrderService) Delete(orderID int) (int, error) {
	rowsAffected, err := o.orderRepository.Delete(orderID)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
