package auth

import (
	"errors"
	"maksehat/data"
	"maksehat/internal/model"
	"maksehat/internal/util"
	"time"
)

func Login(username, password string) (*model.User, error) {
	users, err := data.LoadUserData()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(users); i++ {
		if users[i].Username == username && users[i].Password == password {
			return &users[i], nil
		}
	}
	return nil, errors.New("login gagal, username atau password salah")
}

func Register(name, gender, username, password string, dateOfBirth time.Time) error {
	users, err := data.LoadUserData()
	if err != nil {
		return err
	}

	for i := 0; i < len(users); i++ {
		if username == users[i].Username {
			return errors.New("username sudah digunakan")
		}
	}

	userID := util.GenerateUserID(gender)

	newUser := model.User{
		UserID: userID,
		Name: name,
		Gender: gender,
		DateOfBirth: dateOfBirth,
		Username: username,
		Password: password,
		IsAdmin: false,
	}
	users = append(users, newUser)
	err = data.SaveUserData(users)
	if err != nil {
		return err
	}
	return nil
}
