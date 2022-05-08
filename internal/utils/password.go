package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(password_input, password_storage string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password_storage), []byte(password_input))
	if err != nil {
		return false
	} else {
		return true
	}
}

func HashPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
