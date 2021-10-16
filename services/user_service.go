package services

import (
	"awesomeProject/models"
	"awesomeProject/repository"
	"golang.org/x/crypto/bcrypt"
)

func NewUserService(userRepository repository.UserRepositoryI) *UserService {
	return &UserService{
		userRepository,
	}
}

type UserServiceI interface {
	Create(user *models.User) (int, error)
	GetByID(userID int) (*models.UserResponse, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) (int, error)
}

type UserService struct {
	userRepository repository.UserRepositoryI
}

func (u UserService) Create(user *models.User) (int, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	user.PasswordHash = string(hashedPassword)
	lastID, err := u.userRepository.Create(user)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (u UserService) GetByID(userID int) (*models.UserResponse, error) {
	user, err := u.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	userResponse := models.UserModelTransform(user)
	return userResponse, nil
}

func (u UserService) GetByEmail(email string) (*models.User, error) {
	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) Update(user *models.User) (int, error) {
	rowsAffected, err := u.userRepository.Update(user)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
