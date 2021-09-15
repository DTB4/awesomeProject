package repository

import (
	"awesomeProject/models"
	"database/sql"
)

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

type TokenRepositoryI interface {
	GetByUID(uID string) (int, error)
	Update(userID int, uID string) (sql.Result, error)
	Delete(userID string) (sql.Result, error)
}

type TokenRepository struct {
	db *sql.DB
}

func (t TokenRepository) GetByUID(uID string) (int, error) {
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

func (t TokenRepository) Update(userID int, uID string) (sql.Result, error) {
	result, err := t.db.Exec("UPDATE uids SET user_id=?, uid=?", userID, uID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t TokenRepository) Delete(userID string) (sql.Result, error) {
	result, err := t.db.Exec("DELETE FROM uids WHERE user_id=?", userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
