package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
	"golang.org/x/crypto/bcrypt"
)

func NewUserService(OrderRepository *repository.OrderRepository, UserRepository *repository.UserRepository) *UserService {
	return &UserService{
		*OrderRepository,
		*UserRepository,
	}
}

type UserServiceI interface {
	CreateNewUser(user *models.User) error
	GetUserByID(userID int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	EditUserProfile(user *models.User) error
	GetAllUsers() (*[]models.User, error)
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

func (u UserService) GetUserByID(userID int) (*models.User, error) {
	user, err := u.UserRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := u.UserRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) EditUserProfile(user *models.User) error {
	err := u.UserRepository.EditUserData(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) GetAllUsers() (*[]models.User, error) {
	users, err := u.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
