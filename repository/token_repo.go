package repository

import (
	"awesomeProject/models"
	"database/sql"
	"errors"
)

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

type TokenRepositoryI interface {
	CreateUIDRow(userID int) (int, error)
	GetByUID(uID string) (int, error)
	Update(userID int, uID string) (int, error)
	NullUID(userID int) (int, error)
}

type TokenRepository struct {
	db *sql.DB
}

func (t TokenRepository) CreateUIDRow(userID int) (int, error) {
	result, err := t.db.Exec("INSERT INTO uids (user_id) VALUES (?)", userID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowsAffected == 0 {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (t TokenRepository) GetByUID(uID string) (int, error) {
	if uID == "" {
		return 0, errors.New("empty uID string")
	}
	tokensIDs := models.TokenIDs{}
	rows, err := t.db.Query("SELECT user_id FROM uids WHERE uid=?", uID)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&tokensIDs.UserID)
		if err != nil {
			return 0, err
		}
	}
	err = rows.Close()
	if err != nil {
		return 0, err
	}

	return tokensIDs.UserID, nil
}

func (t TokenRepository) Update(userID int, uID string) (int, error) {
	if userID == 0 {
		return 0, errors.New("user ID is 0")
	}
	if uID == "" {
		return 0, errors.New("uID is empty")
	}
	result, err := t.db.Exec("UPDATE uids SET uid=? WHERE user_id=?", uID, userID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (t TokenRepository) NullUID(userID int) (int, error) {
	if userID == 0 {
		return 0, errors.New("user ID is 0")
	}
	result, err := t.db.Exec("UPDATE uids SET uid=NULL WHERE user_id=?", userID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}
