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
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateNewUser(user *models.User) (sql.Result, error)
	EditUserData(user *models.User) (sql.Result, error)
	GetAllUsers() (*[]models.User, error)
	DeleteUser(id int) (sql.Result, error)
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
		//fmt.Println(user)
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
			//log.Println(err)
		}
		//fmt.Println(user)
	}
	return &user, nil
}

func (u UserRepository) CreateNewUser(user *models.User) (sql.Result, error) {
	result, err := u.db.Exec("INSERT INTO users (first_name, second_name, email, password_hash, created) VALUES (?, ?, ?, ?, ?)", user.FirstName, user.SecondName, user.Email, user.PasswordHash, time.Now())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u UserRepository) EditUserData(user *models.User) (sql.Result, error) {
	result, err := u.db.Exec("UPDATE users SET first_name = ?, second_name = ?, email=?, password_hash=?, updated=?, deleted=? WHERE id=?", user.FirstName, user.SecondName, user.Email, user.PasswordHash, time.Now(), user.Deleted, user.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u UserRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	user := models.User{}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Email, &user.PasswordHash, &user.Created, &user.Updated, &user.Deleted)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}
	return &users, nil
}

func (u UserRepository) DeleteUser(id int) (sql.Result, error) {
	result, err := u.db.Exec("DELETE from users WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return result, nil

}
