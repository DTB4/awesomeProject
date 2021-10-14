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

func NewSupplierHandler(supplierService services.SupplierServiceI, logger *logger.Logger) *SupplierHandler {
	return &SupplierHandler{
		supplierService: supplierService,
		logger:          logger,
	}
}

type SupplierHandlerI interface {
	GetSupplierByID(w http.ResponseWriter, req *http.Request)
	GetAllSuppliers(w http.ResponseWriter, req *http.Request)
	GetByParams(w http.ResponseWriter, req *http.Request)
	GetSuppliersByType(w http.ResponseWriter, req *http.Request)
	GetSuppliersByTime(w http.ResponseWriter, req *http.Request)
}

type SupplierHandler struct {
	supplierService services.SupplierServiceI
	logger          *logger.Logger
}

func (h SupplierHandler) GetSupplierByID(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":

		reqSupplier := new(models.SupplierRequestID)
		err := json.NewDecoder(req.Body).Decode(&reqSupplier)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		supplier, err := h.supplierService.GetByID(reqSupplier.ID)

		if supplier.ID == 0 {
			http.Error(w, "no such supplier", http.StatusNotAcceptable)
			return
		}

		jSupplier, err := json.Marshal(*supplier)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jSupplier)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (h SupplierHandler) GetAllSuppliers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

		suppliers, err := h.supplierService.GetAll()

		jSuppliers, err := json.Marshal(*suppliers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		length, err := w.Write(jSuppliers)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (h SupplierHandler) GetByParams(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

		query := req.URL.Query()
		time := query.Get("_time")
		sType := query.Get("_type")

		suppliers, err := h.supplierService.GetByParams(sType, time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		jSuppliers, err := json.Marshal(*suppliers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jSuppliers)

		if err != nil {
			h.logger.ErrorLog("response error:", err)
		}
		h.logger.InfoLog("GetByParams handler response length: ", length)

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (h SupplierHandler) GetSuppliersByType(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":

		reqSupplier := new(models.SupplierRequestType)
		err := json.NewDecoder(req.Body).Decode(&reqSupplier)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		suppliers, err := h.supplierService.GetAllByType(reqSupplier.Type)

		jSuppliers, err := json.Marshal(*suppliers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jSuppliers)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}
func (h SupplierHandler) GetSuppliersByTime(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":

		reqSupplier := new(models.SupplierRequestTime)
		err := json.NewDecoder(req.Body).Decode(&reqSupplier)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		suppliers, err := h.supplierService.GetAllByTime(reqSupplier.Time)

		jSuppliers, err := json.Marshal(*suppliers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jSuppliers)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (h SupplierHandler) GetSuppliersTypes(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

		types, err := h.supplierService.GetUniqTypes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jTypes, err := json.Marshal(types)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jTypes)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)
		//TODO: replace all fmt.Println with custom logger with useful [info] message!
	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}
