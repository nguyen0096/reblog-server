package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 31)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
