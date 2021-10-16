package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
	"fmt"
	"math"
)

func NewOrderService(orderRepository repository.OrderRepositoryI, orderProductsRepository repository.OrderProductsRepositoryI) *OrderService {
	return &OrderService{
		orderRepository,
		orderProductsRepository,
	}
}

type OrderServiceI interface {
	CreateOrder(order *models.Order) (int, error)
	CreateOrderProducts(orderID int, orderProducts *[]models.OrderProduct) (int, float64, error)
	GetByID(orderID int) (*[]models.OrderProduct, error)
	GetAll(userID int) (*[]models.Order, error)
	Update(updatedOrder *models.Order) (int, error)
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

func (o OrderService) CreateOrderProducts(orderID int, orderProducts *[]models.OrderProduct) (int, float64, error) {
	var rowsAffectedTotal int
	var total float64
	for i := range *orderProducts {
		actualPrice, err := o.orderProductsRepository.GetProductPrice((*orderProducts)[i].ProductID)

		total = math.Round(total*100+float64(int(actualPrice*100)*(*orderProducts)[i].Quantity)) / 100

		fmt.Println(total)
		if err != nil {
			return 0, 0, err
		}
		(*orderProducts)[i].Price = actualPrice
	}

	for i := range *orderProducts {
		(*orderProducts)[i].OrderID = orderID
		rowsAffected, err := o.orderProductsRepository.Create(&(*orderProducts)[i])
		if err != nil {
			return 0, 0, err
		}
		rowsAffectedTotal += rowsAffected
	}

	return rowsAffectedTotal, total, nil
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

func (o OrderService) Update(updatedOrder *models.Order) (int, error) {
	if updatedOrder.Total != 0 {
		rowsAffected, err := o.orderRepository.SetTotal(updatedOrder.ID, updatedOrder.Total)
		if err != nil {
			return 0, err
		}
		return rowsAffected, nil
	} else {
		rowsAffected, err := o.orderRepository.UpdateStatus(updatedOrder)
		if err != nil {
			return 0, err
		}
		return rowsAffected, nil
	}
}

func (o OrderService) Delete(orderID int) (int, error) {
	rowsAffected, err := o.orderRepository.Delete(orderID)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
