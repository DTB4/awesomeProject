package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
	UID string `json:"u_id"`
}

func NewTokenService(cfg *models.AuthConfig, tokenRepository repository.TokenRepositoryI) *TokenService {
	return &TokenService{
		cfg:             cfg,
		tokenRepository: tokenRepository,
	}
}

type TokenServiceI interface {
	ValidateAccessToken(tokenString string) (*JwtCustomClaims, error)
	ValidateRefreshToken(tokenString string) (*JwtCustomClaims, error)
	validateToken(tokenString, secret string) (*JwtCustomClaims, error)
	ParseFromBearerString(input string) string
	generateToken(userID, lifeTimeMinutes int, secret string) (string, string, error)
	GeneratePairOfTokens(userID int) (string, string, error)
	CheckUID(uID string) (int, error)
}

type TokenService struct {
	cfg             *models.AuthConfig
	tokenRepository repository.TokenRepositoryI
}

func (t TokenService) ValidateAccessToken(tokenString string) (*JwtCustomClaims, error) {
	claims, err := t.validateToken(tokenString, t.cfg.AccessSecretString)
	return claims, err
}

func (t TokenService) ValidateRefreshToken(tokenString string) (*JwtCustomClaims, error) {
	claims, err := t.validateToken(tokenString, t.cfg.RefreshSecretString)
	return claims, err
}

func (t TokenService) validateToken(tokenString, secret string) (*JwtCustomClaims, error) {
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

func (t TokenService) generateToken(userID, lifeTimeMinutes int, secret string) (string, string, error) {
	uID := uuid.New().String()
	claims := &JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifeTimeMinutes)).Unix(),
		},
		uID,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, uID, err
}

func (t TokenService) GeneratePairOfTokens(userID int) (string, string, error) {
	accessToken, uid, err := t.generateToken(userID, t.cfg.AccessLifeTimeMinutes, t.cfg.AccessSecretString)
	if err != nil {
		return "", "", err
	}
	refreshToken, _, err := t.generateToken(userID, t.cfg.RefreshLifeTimeMinutes, t.cfg.RefreshSecretString)
	if err != nil {
		return "", "", err
	}

	result, err := t.tokenRepository.Update(userID, uid)
	if err != nil || result == nil {
		return "", "", errors.New("fail to update uid")
	}

	return accessToken, refreshToken, nil
}

func (t TokenService) CheckUID(uID string) (int, error) {
	userID, err := t.tokenRepository.GetByUID(uID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
