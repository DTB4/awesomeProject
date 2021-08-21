package handlers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
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
	ShowUserProfile(w http.ResponseWriter, req *http.Request)
	EditUserProfile(w http.ResponseWriter, req *http.Request)
	TokenCheck(next http.HandlerFunc) http.HandlerFunc
	Login(w http.ResponseWriter, req *http.Request)
}

type ProfileHandler struct {
	userService  *services.UserService
	tokenService *services.TokenService
}

func (p ProfileHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

		users, err := p.userService.GetAllUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		jUsers, err := json.Marshal(*users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jUsers)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (p ProfileHandler) ShowUserProfile(w http.ResponseWriter, req *http.Request) {

	userID := req.Context().Value("CurrentUser").(models.ActiveUserData).ID

	user, err := p.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jUser, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	length, err := w.Write(jUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(length)
}

func (p ProfileHandler) Login(w http.ResponseWriter, req *http.Request) {
	loginForm := new(models.LoginForm)
	err := json.NewDecoder(req.Body).Decode(&loginForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	user, err := p.userService.GetUserByEmail(loginForm.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginForm.Password))
	if err != nil {
		http.Error(w, "invalid input", http.StatusUnauthorized)
		return
	}
	accessString, err := p.tokenService.GenerateToken(user.ID, viper.GetInt("ACCESS_LIFE_TIME_MINUTES"), viper.GetString("ACCESS_SECRET_STRING"))
	refreshString, err := p.tokenService.GenerateToken(user.ID, viper.GetInt("REFRESH_LIFE_TIME_MINUTES"), viper.GetString("REFRESH_SECRET_STRING"))
	if err != nil {
		http.Error(w, "Fail to generate tokens", http.StatusUnauthorized)
	}

	resp := &models.TokenPair{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}
	respJ, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	length, err := w.Write(respJ)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(length)

}

func (p ProfileHandler) EditUserProfile(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		user := new(models.User)
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
		}
		err = p.userService.EditUserProfile(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
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
		curUser := models.ActiveUserData{
			ID: claims.ID,
		}
		req = req.WithContext(context.WithValue(req.Context(), "CurrentUser", curUser))
		next(w, req)
	}
}
