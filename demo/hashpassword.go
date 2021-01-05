package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password []byte) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func main() {
	password := []byte("123")

	res := hashPassword(password)

	log.Println("Password: ", res)

}
