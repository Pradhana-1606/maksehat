package auth

import (
	"errors"
	"strings"
	"unicode"
)

func UsernameValidator(username string) error {
	if len(username) < 3 || len(username) > 26 {
		return errors.New("panjang username harus diantara 3-26 karakter")
	} else if strings.Contains(username, " ") {
		return errors.New("username tidak boleh mengandung spasi")
	} else {
		runeString := []rune(username)
		for i := 0; i < len(runeString); i++ {
			if !unicode.IsLower(runeString[i]) && !unicode.IsDigit(runeString[i]) {
				return errors.New("username harus huruf kecil dan tidak boleh mengandung karakter khusus")
			}
		}
	}
	return nil
}

func PasswordValidator(password string) error {
	if len(password) < 8 {
		return errors.New("password harus lebih dari 8 karakter")
	} else {
		return nil
	}
}
