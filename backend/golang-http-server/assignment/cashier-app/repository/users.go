package repository

import (
	"fmt"
	"strconv"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	records, err := u.db.Load("users")
	if err != nil {
		records = [][]string{
			{"username", "password", "loggedin"},
		}
		if err := u.db.Save("users", records); err != nil {
			return nil, err
		}
	}

	result := make([]User, 0)
	for i := 1; i < len(records); i++ {
		loggedin, err := strconv.ParseBool(records[i][2])
		if err != nil {
			return nil, err
		}

		user := User{
			Username: records[i][0],
			Password: records[i][1],
			Loggedin: loggedin,
			Role:     records[i][3],
		}
		result = append(result, user)
	}
	return result, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	users, err := u.SelectAll()
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

func (u *UserRepository) Save(users []User) error {
	records := [][]string{
		{"username", "password", "loggedin", "role"},
	}

	for _, user := range users {
		records = append(records, []string{
			user.Username,
			user.Password,
			strconv.FormatBool(user.Loggedin),
			user.Role,
		})
	}
	return u.db.Save("users", records)
}

func (u *UserRepository) GetUserRole(username string) (*string, error) {
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			return &user.Username, nil
		}
	}

	return nil, fmt.Errorf("user role not found")
}
