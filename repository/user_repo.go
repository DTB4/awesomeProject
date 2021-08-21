package repository

import (
	"awesomeProject/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepositoryI interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateNewUser(user *models.User) error
	EditUserData(user *models.User) error
	GetAll()
}

type UserRepository struct {
	db *sql.DB
}

func (u UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	rows, err := u.db.Query("SELECT * FROM users WHERE email=?", email)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Email, &user.PasswordHash, &user.Created, &user.Updated, &user.Deleted)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(user)
	}
	return &user, nil
}

func (u UserRepository) GetUserByID(userID int) (*models.User, error) {
	user := models.User{}
	rows, err := u.db.Query("SELECT * FROM users WHERE id=?", userID)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Email, &user.PasswordHash, &user.Created, &user.Updated, &user.Deleted)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(user)
	}
	return &user, nil
}

func (u UserRepository) CreateNewUser(user *models.User) error {
	result, err := u.db.Exec("INSERT INTO users (first_name, second_name, email, password_hash, created) VALUES (?, ?, ?, ?, ?)", user.FirstName, user.SecondName, user.Email, user.PasswordHash, time.Now())
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

func (u UserRepository) EditUserData(user *models.User) error {
	result, err := u.db.Exec("UPDATE users SET first_name = ?, second_name = ?, email=?, password_hash=?, updated=? WHERE id=?", user.FirstName, user.SecondName, user.Email, user.PasswordHash, time.Now(), user.ID)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

func (u UserRepository) GetAll() {
	var users []models.User
	rows, err := u.db.Query("SELECT * FROM users?")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; rows.Next(); i++ {
		err := rows.Scan(&users[i].ID, &users[i].FirstName, &users[i].SecondName, &users[i].Email, &users[i].PasswordHash, &users[i].Created, &users[i].Updated, &users[i].Deleted)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(users[i])
	}
	//return &users[], nil
}
