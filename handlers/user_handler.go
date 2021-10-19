package handlers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"context"
	"encoding/json"
	"github.com/DTB4/logger/v2"
	"golang.org/x/crypto/bcrypt"
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
		p.logger.FErrorLog("Error in CreateNewUser handler while decoding request:", err.Error())
		return
	}
	lastUserId, err := p.userService.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		p.logger.FErrorLog("Error in CreateNewUser handler while creating user:", err.Error())
		return
	}
	err = p.tokenService.CreateUIDRow(lastUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		p.logger.FErrorLog("Error in CreateNewUser handler while creating UID row in DB:", err.Error())
		return
	}
	userResponse := models.UserResponse{
		ID:        lastUserId,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		p.logger.FErrorLog("Error in CreateNewUser handler while coding response in json:", err.Error())
		return
	}
	length, err := w.Write(response)
	if err != nil || length == 0 {
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
		p.logger.FErrorLog("Error in CreateNewUser handler while writing response:", err.Error())
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
			p.logger.FErrorLog("Error in ShowUserProfile handler while get user form DB:", err.Error())
			return
		}
		jUser, err := json.Marshal(*user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			p.logger.FErrorLog("Error in ShowUserProfile handler while coding response in json:", err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jUser)
		if err != nil {
			p.logger.FErrorLog("Error in ShowUserProfile handler while writing response:", err.Error())
		}
		p.logger.FInfoLog("ShowUserProfile responded with length", length)
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
			p.logger.FErrorLog("Error in Login handler while decoding request:", err.Error())
			return
		}
		user, err := p.userService.GetByEmail(loginForm.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			p.logger.FErrorLog("Error in Login handler while get user from DB:", err.Error())
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginForm.Password))
		if err != nil {
			http.Error(w, "invalid input", http.StatusUnauthorized)
			p.logger.FErrorLog("Error in Login handler while comparing password hashes:", err.Error())
			return
		}
		accessString, refreshString, err := p.tokenService.GeneratePairOfTokens(user.ID)
		if err != nil {
			http.Error(w, "Fail to generate tokens", http.StatusUnauthorized)
			p.logger.FErrorLog("Error in Login handler while generating tokens:", err.Error())
			return
		}

		resp := &models.TokenPair{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}
		respJ, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusOK)
		length, err := w.Write(respJ)
		if err != nil {
			p.logger.FErrorLog("Error in Login handler while writing response:", err.Error())
			return
		}
		p.logger.FInfoLog("Login handler responded with length", length)

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
			p.logger.FErrorLog("Error in Refresh handler while generating tokens:", err.Error())
			return
		}

		resp := &models.TokenPair{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}
		respJ, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusOK)
		length, err := w.Write(respJ)
		if err != nil {
			p.logger.FErrorLog("Error in Refresh handler while writing response:", err.Error())
			return
		}
		p.logger.FInfoLog("Refresh handler responded with length", length)
	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (p UserHandler) Logout(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

		userID := req.Context().Value("CurrentUser").(models.ActiveUserData).ID
		err := p.tokenService.Logout(userID)
		if err != nil {
			http.Error(w, "fail to logout current user ", http.StatusMethodNotAllowed)
			p.logger.FErrorLog("Error in Logout handler:", err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}
