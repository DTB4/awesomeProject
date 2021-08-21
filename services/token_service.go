package services

import (
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

type TokenServiceI interface {
	ValidateToken(tokenString, secret string) (*JwtCustomClaims, error)
	GetTokenFromBearerString(input string) string
	GenerateToken(userID, lifeTimeMinutes int, secret string) (string, error)
	GenerateAccessToken(id int) (string, error)
	GenerateRefreshToken(id int) (string, error)
}

type TokenService struct {
}

func (t TokenService) ValidateToken(tokenString, secret string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func (t TokenService) GetTokenFromBearerString(input string) string {
	if input == "" {
		return ""
	}
	parts := strings.Split(input, "Bearer")
	if len(parts) != 2 {
		return ""
	}
	token := strings.TrimSpace(parts[1])
	if len(token) == 0 {
		return ""
	}
	return token
}

func (t TokenService) GenerateToken(userID, lifeTimeMinutes int, secret string) (string, error) {
	claims := &JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifeTimeMinutes)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
