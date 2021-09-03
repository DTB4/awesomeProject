package repository

import (
	"awesomeProject/models"
	"database/sql"
	"log"
	"time"
)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepositoryI interface {
	Create(user *models.User) (sql.Result, error)
	GetByEmail(email string) (*models.User, error)
	GetByID(id int) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(user *models.User) (sql.Result, error)
	Delete(id int) (sql.Result, error)
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
		err = rows.Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Email, &user.PasswordHash, &user.Created, &user.Updated, &user.Deleted)
		if err != nil {
			log.Println(err)
		}
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
		err = rows.Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Email, &user.PasswordHash, &user.Created, &user.Updated, &user.Deleted)
		if err != nil {
			log.Println(err)
		}
	}
	return &user, nil
}

func (u UserRepository) Create(user *models.User) (sql.Result, error) {
	result, err := u.db.Exec("INSERT INTO users (first_name, second_name, email, password_hash, created) VALUES (?, ?, ?, ?, ?)", user.FirstName, user.SecondName, user.Email, user.PasswordHash, time.Now())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u UserRepository) Update(user *models.User) (sql.Result, error) {
	result, err := u.db.Exec("UPDATE users SET first_name = ?, second_name = ?, email=?, password_hash=?, updated=? WHERE id=? AND deleted=false", user.FirstName, user.SecondName, user.Email, user.PasswordHash, time.Now(), user.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u UserRepository) GetAll() (*[]models.User, error) {
	var users []models.User
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	user := models.User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Email, &user.PasswordHash, &user.Created, &user.Updated, &user.Deleted)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u UserRepository) Delete(id int) (sql.Result, error) {
	result, err := u.db.Exec("DELETE from users WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
