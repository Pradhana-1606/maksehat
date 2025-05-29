package service

import (
	"fmt"
	"maksehat/data"
	"maksehat/internal/model"
	"time"
)

func IsAdminExist() error {
	users, err := data.LoadUserData()
	if err != nil {
		return err
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == "admin" {
			return nil
		}
	}

	admin := model.User{
		UserID: "0000000000",
		Name: "admin",
		Gender: "none",
		DateOfBirth: time.Date(2025, 05, 28, 21, 40, 0, 0, time.UTC),
		Username: "admin",
		Password: "admin",
		IsAdmin: true,
	}
	users = append(users, admin)
	err = data.SaveUserData(users)
	if err != nil {
		return err
	}
	return nil
}

func GetName(userID string) (string, error) {
	users, err := data.LoadUserData()
	if err != nil {
		return "", fmt.Errorf("gagal memuat data pengguna: %v", err)
	}
	for i := 0; i < len(users); i++ {
		if userID == users[i].UserID {
			return users[i].Name, nil
		}
	}
	return "", fmt.Errorf("pengguna dengan ID %s tidak ditemukan", userID)
}