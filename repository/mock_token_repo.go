package repository

import (
	"errors"
	"github.com/google/uuid"
	"log"
)

func NewMockTokenRepository() *MockTokenRepository {
	return &MockTokenRepository{}
}

type MockTokenRepository struct {
}

func (m MockTokenRepository) GetByUID(uID string) (int, error) {
	if uID != "" {
		uuidVal, err := uuid.Parse(uID)
		if err != nil {
			return 1, nil
		}
		log.Println(uuidVal)
	}
	return 0, errors.New("uid not found")
}

func (m MockTokenRepository) Update(userID int, uID string) (int, error) {
	if uID != "" && userID != 0 {
		uuidVal, err := uuid.Parse(uID)
		if err != nil {
			return 1, nil
		}
		log.Println(uuidVal)
	}
	if userID == 0 {
		return 0, errors.New("0 user id")
	}
	if uID == "" {
		return 0, errors.New("empty uID")
	}
	return 0, errors.New("nothing updated")
}

//NullUID 0- user is empty, 1 for success, 2 for user not found, any other - nothing updated
func (m MockTokenRepository) NullUID(userID int) (int, error) {
	if userID == 0 {
		return 0, errors.New("user is empty")
	}
	if userID == 1 {
		return 1, nil
	}
	if userID == 2 {
		return 0, errors.New("user not found")
	}
	return 0, errors.New("nothing updated")
}
