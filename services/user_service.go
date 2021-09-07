package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

func NewUserService(UserRepository *repository.UserRepository) *UserService {
	return &UserService{
		*UserRepository,
	}
}

type UserServiceI interface {
	Create(user *models.User) (sql.Result, error)
	GetByID(userID int) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) (sql.Result, error)
	GetAll() (*[]models.User, error)
}

type UserService struct {
	UserRepository repository.UserRepository
}

func (u UserService) CreateNewUser(user *models.User) (sql.Result, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	user.PasswordHash = string(hashedPassword)
	result, err := u.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u UserService) GetUserByID(userID int) (*models.User, error) {
	user, err := u.UserRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := u.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) EditUserProfile(user *models.User) (sql.Result, error) {
	result, err := u.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u UserService) GetAllUsers() (*[]models.User, error) {
	users, err := u.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
