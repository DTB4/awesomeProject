package services

import (
	"awesomeProject/models"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func NewTokenService(cfg *models.AuthConfig) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

type TokenServiceI interface {
	Validate(tokenString, secret string) (*JwtCustomClaims, error)
	ParseFromBearerString(input string) string
	generateToken(userID, lifeTimeMinutes int, secret string) (string, error)
	GeneratePairOfTokens(userID int) (string, string, error)
}

type TokenService struct {
	cfg *models.AuthConfig
}

func (t TokenService) Validate(tokenString, secret string) (*JwtCustomClaims, error) {
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

func (t TokenService) ParseFromBearerString(input string) string {
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

func (t TokenService) generateToken(userID, lifeTimeMinutes int, secret string) (string, error) {
	claims := &JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifeTimeMinutes)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (t TokenService) GeneratePairOfTokens(userID int) (string, string, error) {
	accessToken, err := t.generateToken(userID, t.cfg.AccessLifeTimeMinutes, t.cfg.AccessSecretString)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := t.generateToken(userID, t.cfg.RefreshLifeTimeMinutes, t.cfg.RefreshSecretString)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
