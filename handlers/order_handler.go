package handlers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"encoding/json"
	"github.com/DTB4/logger/v2"
	"net/http"
)

func NewOrderHandler(orderService services.OrderServiceI, logger *logger.Logger) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
		logger:       logger,
	}
}

type OrderHandlerI interface {
	Create(w http.ResponseWriter, req *http.Request)
	GetByID(w http.ResponseWriter, req *http.Request)
	GetAll(w http.ResponseWriter, req *http.Request)
	Update(w http.ResponseWriter, req *http.Request)
	Delete(w http.ResponseWriter, req *http.Request)
}

type OrderHandler struct {
	orderService services.OrderServiceI
	logger       *logger.Logger
}

func (o OrderHandler) Create(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":

		var orderProducts []models.OrderProduct
		err := json.NewDecoder(req.Body).Decode(&orderProducts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
		}
		order := models.Order{
			IDUser: req.Context().Value("CurrentUser").(models.ActiveUserData).ID,
			Status: "created",
		}
		orderCreationResult, err := o.orderService.CreateOrder(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orderID, err := orderCreationResult.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orderProductsCreationResult, err := o.orderService.CreateOrderProducts(int(orderID), &orderProducts)
		if err != nil || orderProductsCreationResult == 0 {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (o OrderHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		requesterID := new(models.RequestOrderID)
		err := json.NewDecoder(req.Body).Decode(&requesterID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		orderProducts, err := o.orderService.GetByID(requesterID.OrderId)

		respJ, err := json.Marshal(orderProducts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(respJ)
		if err != nil || length == 0 {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (o OrderHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		userID := req.Context().Value("CurrentUser").(models.ActiveUserData).ID
		orders, err := o.orderService.GetAll(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		respJ, err := json.Marshal(orders)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(respJ)
		if err != nil || length == 0 {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (o OrderHandler) Update(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		updateOrderRequest := new(models.UpdateOrderRequest)
		err := json.NewDecoder(req.Body).Decode(&updateOrderRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result, err := o.orderService.Update(updateOrderRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if rowsAffected == 0 {
			http.Error(w, "nothing was changed", http.StatusNotModified)
			return
		}

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (o OrderHandler) Delete(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}
