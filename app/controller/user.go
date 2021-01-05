package controller

import (
	"database/sql"
	"log"
	"reblog-server/dto"

	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	base *baseController
}

func newUserController(base baseController) *userController {
	return &userController{
		base: &base,
	}
}

func (c *userController) CreateUser(form dto.LoginForm) error {

	// Check if user existing
	_, err := c.base.store.User().Get(form.Username)

	if err == sql.ErrNoRows {

	}

	x, err := bcrypt.GenerateFromPassword([]byte("test password"), 4)
	if err != nil {
		log.Println("Failed to generate hashing from password")
	}
	log.Println(string(x))

	return nil
}
