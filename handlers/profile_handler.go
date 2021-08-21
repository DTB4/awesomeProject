package handlers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

func NewProfileHandler(userService *services.UserService, tokenService *services.TokenService) *ProfileHandler {
	return &ProfileHandler{
		userService:  userService,
		tokenService: tokenService,
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
	userService  *services.UserService
	tokenService *services.TokenService
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
	return func(w http.ResponseWriter, req *http.Request) {
		bearerString := req.Header.Get("Authorization")
		tokenString := p.tokenService.GetTokenFromBearerString(bearerString)

		claims, err := p.tokenService.ValidateToken(tokenString, viper.GetString("ACCESS_SECRET_STRING"))

		fmt.Println("user with ID = ", claims.ID, " login")

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next(w, req)
	}
}
