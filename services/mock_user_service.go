package services

import (
	"awesomeProject/models"
	"errors"
)

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

type MockUserService struct {
}

func (m MockUserService) Create(user *models.User) (int, error) {
	if user != nil {
		return 1, nil
	}
	return 0, errors.New("user is empty")
}

func (m MockUserService) GetByID(userID int) (*models.User, error) {
	if userID != 0 {
		user := models.User{
			ID:        userID,
			Email:     "email",
			FirstName: "FirstName",
			LastName:  "LastName",
		}
		return &user, nil
	}
	return nil, errors.New("user ID is 0")
}

func (m MockUserService) GetByEmail(email string) (*models.User, error) {
	if email != "" {
		return &models.User{
			ID:        1,
			Email:     email,
			FirstName: "FirstName",
			LastName:  "LastName",
		}, nil
	}
	return nil, errors.New("user email is empty")
}

func (m MockUserService) Update(user *models.User) (int, error) {
	if user != nil {
		return 1, nil
	}
	return 0, errors.New("nothing changed")
}
