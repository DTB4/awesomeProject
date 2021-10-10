package midleware

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"context"
	"github.com/DTB4/logger/v2"
	"net/http"
)

func NewAuthHandler(tokenService services.TokenServiceI, logger *logger.Logger) *AuthHandler {
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
	tokenService services.TokenServiceI
	logger       *logger.Logger
}

func (ah AuthHandler) AccessTokenCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		bearerString := req.Header.Get("Authorization")
		tokenString := ah.tokenService.ParseFromBearerString(bearerString)
		claims, err := ah.tokenService.ValidateAccessToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		userFromDB, err := ah.tokenService.CheckUID(claims.UID)
		if err != nil || claims.ID != userFromDB {
			http.Error(w, "logout", http.StatusUnauthorized)
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
		bearerString := req.Header.Get("Authorization")
		tokenString := ah.tokenService.ParseFromBearerString(bearerString)
		claims, err := ah.tokenService.ValidateRefreshToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		userFromDB, err := ah.tokenService.CheckUID(claims.UID)
		if err != nil || claims.ID != userFromDB {
			http.Error(w, "logout", http.StatusUnauthorized)
			return
		}
		curUser := models.ActiveUserData{
			ID: claims.ID,
		}

		req = req.WithContext(context.WithValue(req.Context(), "CurrentUser", curUser))
		next(w, req)
	}
}
