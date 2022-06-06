package repository

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (User, error) {
	var user User

	rows, err := u.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	var users []User

	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) Login(username string, password string) (*string, error) {
	users, err := u.FetchUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user.Username, nil
		}
	}

	return nil, fmt.Errorf("Login Failed")
}

func (u *UserRepository) InsertUser(username string, password string, role string, loggedin bool) error {
	_, err := u.db.Exec("INSERT INTO users (username, password, role, loggedin) VALUES (?, ?, ?, ?)", username, password, role, loggedin)
	return err
}

func (u *UserRepository) FetchUserRole(username string) (*string, error) {
	var role string

	rows, err := u.db.Query("SELECT role FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&role)
		if err != nil {
			return nil, err
		}
	}

	return &role, nil
}
