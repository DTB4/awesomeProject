package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func NewUserService(OrderRepository *repository.OrderRepository, UserRepository *repository.UserRepository) *UserService {
	return &UserService{
		*OrderRepository,
		*UserRepository,
	}
}

type UserServiceI interface {
	CreateNewUser(user *models.User) error
	GetUserByID(userID int) (models.User, error)
	EditUserProfile(user models.User) error
	ValidateToken(tokenString, secret string) (*JwtCustomClaims, error)
	GetTokenFromBearerString(input string) string
	GenerateAccessToken(id int) (string, error)
	GenerateRefreshToken(id int) (string, error)
}

type UserService struct {
	OrderRepository repository.OrderRepository
	UserRepository  repository.UserRepository
}

func (u UserService) CreateNewUser(user *models.User) error {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	user.PasswordHash = string(hashedPassword)
	err := u.UserRepository.CreateNewUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) GetUserByID(userID int) (models.User, error) {
	panic("implement me")
}

func (u UserService) EditUserProfile(user models.User) error {
	panic("implement me")
}

func (u UserService) ValidateToken(tokenString, secret string) (*JwtCustomClaims, error) {
	panic("implement me")
}

func (u UserService) GetTokenFromBearerString(input string) string {
	panic("implement me")
}

func (u UserService) GenerateAccessToken(id int) (string, error) {
	panic("implement me")
}

func (u UserService) GenerateRefreshToken(id int) (string, error) {
	panic("implement me")
}
