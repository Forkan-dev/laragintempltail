package service

import (
	"golang.org/x/crypto/bcrypt"
)

func MakePassword(password string) string {

	bytePassword := []byte(password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
