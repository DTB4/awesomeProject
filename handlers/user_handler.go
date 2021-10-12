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
)

func NewUserHandler(userService services.UserServiceI, tokenService services.TokenServiceI, logger *logger.Logger) *UserHandler {
	return &UserHandler{
		userService:  userService,
		logger:       logger,
		tokenService: tokenService,
	}
}

type UserHandlerI interface {
	CreateNewUser(w http.ResponseWriter, req *http.Request)
	ShowUserProfile(w http.ResponseWriter, req *http.Request)
	EditUserProfile(w http.ResponseWriter, req *http.Request)
	Login(w http.ResponseWriter, req *http.Request)
	Refresh(w http.ResponseWriter, req *http.Request)
	Logout(w http.ResponseWriter, req *http.Request)
}

type UserHandler struct {
	userService  services.UserServiceI
	logger       *logger.Logger
	tokenService services.TokenServiceI
}

func (p UserHandler) CreateNewUser(w http.ResponseWriter, req *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	user := new(models.User)
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
	}
	lastUserId, err := p.userService.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = p.tokenService.CreateUIDRow(lastUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userResponse := models.UserResponse{
		ID:        lastUserId,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Created:   true,
	}
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(userResponse)
	length, err := w.Write(response)
	if err != nil || length == 0 {
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
		return
	}
}

func (p UserHandler) ShowUserProfile(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		userID := req.Context().Value("CurrentUser").(models.ActiveUserData).ID

		user, err := p.userService.GetByID(userID)
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

func (p UserHandler) EditUserProfile(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		user := new(models.User)
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
		}
		rowsAffected, err := p.userService.Update(user)
		if err != nil || rowsAffected == 0 {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		req = req.WithContext(context.WithValue(req.Context(), "CurrentUser", user.ID))
		p.ShowUserProfile(w, req)

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (p UserHandler) Login(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		loginForm := new(models.LoginForm)
		err := json.NewDecoder(req.Body).Decode(&loginForm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		user, err := p.userService.GetByEmail(loginForm.Email)
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

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (p UserHandler) Refresh(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
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
	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (p UserHandler) Logout(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":

		userID := req.Context().Value("CurrentUser").(models.ActiveUserData).ID
		err := p.tokenService.Logout(userID)
		if err != nil {
			http.Error(w, "fail to logout current user ", http.StatusMethodNotAllowed)
			return
		}
		http.RedirectHandler("/index", http.StatusOK)

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}
