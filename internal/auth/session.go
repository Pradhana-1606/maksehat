package auth

import "maksehat/internal/model"

var activeUser *model.User

func SetActiveUser(user *model.User) {
	activeUser = user
}

func GetActiveUser() *model.User {
	return activeUser
}

func IsLoggedIn() bool {
	return activeUser != nil
}

func IsAdmin() bool {
	return activeUser != nil && activeUser.IsAdmin
}

func Logout() {
	activeUser = nil
}