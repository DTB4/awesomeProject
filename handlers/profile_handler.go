package handlers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"context"
	"encoding/json"
	"fmt"
	"github.com/DTB4/logger/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func NewProfileHandler(userService *services.UserService, tokenService *services.TokenService, logger *logger.Logger) *ProfileHandler {
	return &ProfileHandler{
		userService:  userService,
		tokenService: tokenService,
		logger:       logger,
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
	logger       *logger.Logger
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
			p.logger.FatalLog("failed to write jUsers in responseWriter", err)
		}
		fmt.Println(length)

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (p ProfileHandler) ShowUserProfile(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

		userID := req.Context().Value("CurrentUser").(models.ActiveUserData).ID

		user, err := p.userService.GetUserByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		jUser, err := json.Marshal(*user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jUser)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)
	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
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
	accessLifeTimeMinutes, _ := strconv.Atoi(os.Getenv("ACCESS_LIFE_TIME_MINUTES"))
	refreshLifeTimeMinutes, _ := strconv.Atoi(os.Getenv("REFRESH_LIFE_TIME_MINUTES"))
	accessSecretString := os.Getenv("ACCESS_SECRET_STRING")
	refreshSecretString := os.Getenv("REFRESH_SECRET_STRING")
	accessString, err := p.tokenService.GenerateToken(user.ID, accessLifeTimeMinutes, accessSecretString)
	refreshString, err := p.tokenService.GenerateToken(user.ID, refreshLifeTimeMinutes, refreshSecretString)
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
		result, err := p.userService.EditUserProfile(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rowsAffected, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println(rowsAffected)

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
		result, err := p.userService.CreateNewUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		lastUserId, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println(lastUserId)

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (p ProfileHandler) TokenCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		bearerString := req.Header.Get("Authorization")
		tokenString := p.tokenService.GetTokenFromBearerString(bearerString)
		accessSecretString := os.Getenv("ACCESS_SECRET_STRING")
		claims, err := p.tokenService.ValidateToken(tokenString, accessSecretString)

		//fmt.Println("user with ID = ", claims.ID, " login")

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
