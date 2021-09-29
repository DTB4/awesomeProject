package handlers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"encoding/json"
	"fmt"
	"github.com/DTB4/logger/v2"
	"log"
	"net/http"
)

func NewProductHandler(productService services.ProductServiceI, logger *logger.Logger) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		logger:         logger,
	}
}

type ProductHandlerI interface {
	GetProductDyID(w http.ResponseWriter, req *http.Request)
	GetAll(w http.ResponseWriter, req *http.Request)
	GetAllByType(w http.ResponseWriter, req *http.Request)
	GetAllBySupplierID(w http.ResponseWriter, req *http.Request)
}

type ProductHandler struct {
	productService services.ProductServiceI
	logger         *logger.Logger
}

func (p ProductHandler) GetProductDyID(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

		reqProduct := new(models.ProductRequest)
		err := json.NewDecoder(req.Body).Decode(&reqProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		product, err := p.productService.GetByID(reqProduct.ID)

		if product == nil {
			http.Error(w, "no such product", http.StatusNotAcceptable)
			return
		}

		jProduct, err := json.Marshal(*product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jProduct)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (p ProductHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		products, err := p.productService.GetAll()

		jProducts, err := json.Marshal(*products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jProducts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (p ProductHandler) GetAllByType(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

		reqProduct := new(models.ProductTypeRequest)
		err := json.NewDecoder(req.Body).Decode(&reqProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		products, err := p.productService.GetAllByType(reqProduct.ProductType)

		jProducts, err := json.Marshal(*products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jProducts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (p ProductHandler) GetAllBySupplierID(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		reqProduct := new(models.ProductSupplierIDRequest)
		err := json.NewDecoder(req.Body).Decode(&reqProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		products, err := p.productService.GetAllBySuppliersID(reqProduct.SupplierID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		if len(*products) == 0 {
			http.Error(w, "no products available for this supplier", http.StatusNotAcceptable)
			return
		}

		jProducts, err := json.Marshal(*products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jProducts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}
