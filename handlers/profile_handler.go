package handlers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"encoding/json"
	"net/http"
)

func NewProfileHandler(userService *services.UserService) *ProfileHandler {
	return &ProfileHandler{
		userService,
	}
}

type ProfileHandlerI interface {
	GetAll(w http.ResponseWriter, req *http.Request)
	CreateNewUser(w http.ResponseWriter, req *http.Request)
	GetUserByID(w http.ResponseWriter, req *http.Request)
	EditUserProfile(w http.ResponseWriter, req *http.Request)
	TokenCheck(next http.HandlerFunc) http.HandlerFunc
}

type ProfileHandler struct {
	userService *services.UserService
}

func (p ProfileHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	panic("implement me")
}

func (p ProfileHandler) GetUserByID(w http.ResponseWriter, req *http.Request) {
	panic("implement me")
}

func (p ProfileHandler) EditUserProfile(w http.ResponseWriter, req *http.Request) {
	panic("implement me")
}

func (p ProfileHandler) CreateNewUser(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		user := new(models.User)
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
		}
		err = p.userService.CreateNewUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (p ProfileHandler) TokenCheck(next http.HandlerFunc) http.HandlerFunc {
	panic("implement me")
}
