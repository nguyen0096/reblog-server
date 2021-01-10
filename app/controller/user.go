package controller

import (
	"log"
	"reblog-server/config"
	"reblog-server/dto"
	"reblog-server/model"

	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	base *baseController
}

func newUserController(base *baseController) *userController {
	return &userController{
		base: base,
	}
}

func (c *userController) CreateUserFromSignup(form dto.LoginForm) error {
	var hashedPw []byte
	var err error

	if hashedPw, err = bcrypt.GenerateFromPassword([]byte(form.Password), config.App.Controller.HashCost); err != nil {
		log.Println("Failed to generate hashing from password")
		return err
	}

	newUser := model.User{
		Username: form.Username,
		Password: string(hashedPw),
	}

	err = c.base.store.User().Create(newUser)
	return err
}
