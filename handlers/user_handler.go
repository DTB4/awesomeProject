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
	switch req.Method {
	case "POST":
		user := new(models.User)
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
		}
		result, err := p.userService.Create(user)
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

func (p UserHandler) ShowUserProfile(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

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
		result, err := p.userService.Update(user)
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
			http.Error(w, "fail to logout current user", http.StatusMethodNotAllowed)
			return
		}
		http.RedirectHandler("/index", http.StatusOK)

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}
