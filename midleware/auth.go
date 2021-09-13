package midleware

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"context"
	"github.com/DTB4/logger/v2"
	"net/http"
	"os"
)

func NewAuthHandler(tokenService *services.TokenService, logger *logger.Logger) *AuthHandler {
	return &AuthHandler{
		tokenService: tokenService,
		logger:       logger,
	}
}

type AuthHandlerI interface {
	AccessTokenCheck(next http.HandlerFunc) http.HandlerFunc
	RefreshTokenCheck(next http.HandlerFunc) http.HandlerFunc
}

type AuthHandler struct {
	tokenService *services.TokenService
	logger       *logger.Logger
}

func (ah AuthHandler) AccessTokenCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var accessSecretString = os.Getenv("ACCESS_SECRET_STRING")
		bearerString := req.Header.Get("Authorization")
		tokenString := ah.tokenService.ParseFromBearerString(bearerString)
		claims, err := ah.tokenService.Validate(tokenString, accessSecretString)
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

func (ah AuthHandler) RefreshTokenCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var refreshSecretString = os.Getenv("REFRESH_SECRET_STRING")
		bearerString := req.Header.Get("Authorization")
		tokenString := ah.tokenService.ParseFromBearerString(bearerString)
		claims, err := ah.tokenService.Validate(tokenString, refreshSecretString)
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
