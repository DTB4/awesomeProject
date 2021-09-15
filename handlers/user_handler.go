package handlers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"encoding/json"
	"fmt"
	"github.com/DTB4/logger/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func NewProfileHandler(userService *services.UserService, tokenService *services.TokenService, logger *logger.Logger) *ProfileHandler {
	return &ProfileHandler{
		userService:  userService,
		logger:       logger,
		tokenService: tokenService,
	}
}

type ProfileHandlerI interface {
	CreateNewUser(w http.ResponseWriter, req *http.Request)
	ShowUserProfile(w http.ResponseWriter, req *http.Request)
	EditUserProfile(w http.ResponseWriter, req *http.Request)
	Login(w http.ResponseWriter, req *http.Request)
	Refresh(w http.ResponseWriter, req *http.Request)
}

type ProfileHandler struct {
	userService  *services.UserService
	logger       *logger.Logger
	tokenService *services.TokenService
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
	accessString, refreshString, err := p.tokenService.GeneratePairOfTokens(user.ID)
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

func (p ProfileHandler) Refresh(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("CurrentUser").(models.ActiveUserData).ID

	accessString, refreshString, err := p.tokenService.GeneratePairOfTokens(userID)
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
