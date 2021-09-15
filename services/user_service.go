package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

func NewUserService(userRepository repository.UserRepositoryI) *UserService {
	return &UserService{
		userRepository,
	}
}

type UserServiceI interface {
	Create(user *models.User) (sql.Result, error)
	GetByID(userID int) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) (sql.Result, error)
}

type UserService struct {
	userRepository repository.UserRepositoryI
}

func (u UserService) Create(user *models.User) (sql.Result, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	user.PasswordHash = string(hashedPassword)
	result, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u UserService) GetByID(userID int) (*models.User, error) {
	user, err := u.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetByEmail(email string) (*models.User, error) {
	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) Update(user *models.User) (sql.Result, error) {
	result, err := u.userRepository.Update(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
