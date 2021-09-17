package repository

import (
	"awesomeProject/models"
	"database/sql"
	"time"
)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepositoryI interface {
	Create(user *models.User) (int, error)
	GetByEmail(email string) (*models.User, error)
	GetByID(id int) (*models.User, error)
	Update(user *models.User) (int, error)
	Delete(id int) (int, error)
	CreateUIDRow(userID int) (int, error)
}

type UserRepository struct {
	db *sql.DB
}

func (u UserRepository) GetByEmail(email string) (*models.User, error) {
	user := models.User{}
	rows, err := u.db.Query("SELECT * FROM users WHERE email=? AND deleted=false", email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.Created, &user.Updated, &user.Deleted)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) GetByID(id int) (*models.User, error) {
	user := models.User{}
	rows, err := u.db.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.Created, &user.Updated, &user.Deleted)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) Create(user *models.User) (int, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return 0, err
	}
	result, err := tx.Exec("INSERT INTO users (first_name, last_name, email, password_hash, created) VALUES (?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.PasswordHash, time.Now())
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	result2, err := tx.Exec("INSERT INTO uids (user_id) VALUES (?)", user.ID, nil)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	lastID, err := result.LastInsertId()
	rowsAffected, err := result2.RowsAffected()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	if rowsAffected == 0 {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	return int(lastID), nil
}

func (u UserRepository) Update(user *models.User) (int, error) {
	result, err := u.db.Exec("UPDATE users SET first_name = ?, last_name = ?, email=?, password_hash=?, updated=? WHERE id=? AND deleted=false", user.FirstName, user.LastName, user.Email, user.PasswordHash, time.Now(), user.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (u UserRepository) Delete(id int) (int, error) {
	result, err := u.db.Exec("DELETE from users WHERE id=?", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (u UserRepository) CreateUIDRow(userID int) (int, error) {
	result, err := u.db.Exec("INSERT INTO uids (user_id) VALUES (?)", userID, nil)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}
