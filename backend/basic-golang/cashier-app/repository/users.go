package repository

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
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
		return nil, err
	}

	result := make([]User, 0)
	for i := 1; i < len(records); i++ {
		result = append(result, User{
			Username: records[i][0],
			Password: records[i][1],
			Loggedin: false,
		})
	}

	return result, nil

}

func (u *UserRepository) SelectAll() ([]User, error) {
	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}

	return users, nil // TODO: replace this
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	u.LogoutAll()
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == username && users[i].Password == password {
			err := u.changeStatus(username, true)
			if err != nil {
				return nil, err
			}
			return &users[i].Username, nil
		}
	}

	return nil, fmt.Errorf("Login Failed") // TODO: replace this
}

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	records, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	//loop through users
	for i := 0; i < len(records); i++ {
		if records[i].Loggedin {
			return &records[i].Username, nil
		}
	}

	return nil, fmt.Errorf("no user is logged in") // TODO: replace this
}

func (u *UserRepository) Logout(username string) error {
	records, err := u.FindLoggedinUser()
	if err != nil {
		return err
	}
	//if username is logged in, call changeStatus to false
	if *records == username {
		err := u.changeStatus(username, false)
		if err != nil {
			return err
		}
	}
	return nil // TODO: replace this
}

func (u *UserRepository) Save(users []User) error {
	records := [][]string{
		{"username", "password"},
	}

	for i := 0; i < len(users); i++ {
		records = append(records, []string{
			users[i].Username,
			users[i].Password,
		})
	}
	return u.db.Save("users", records)
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	//update
	//change status Loggedin to false
	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			users[i].Loggedin = status
			return u.Save(users)
		}
	}

	return nil // TODO: replace this
}

func (u *UserRepository) LogoutAll() error {
	users, err := u.SelectAll()
	if err != nil {
		return err
	}

	//loop through users
	for i := 0; i < len(users); i++ {
		err := u.Logout(users[i].Username)
		if err != nil {
			return err
		}
	}

	return nil // TODO: replace this
}
