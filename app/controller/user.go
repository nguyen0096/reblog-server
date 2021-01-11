package controller

import (
	"reblog-server/config"
	"reblog-server/model"
	"reblog-server/utils"

	"github.com/google/uuid"
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

func (c *userController) CreateUserFromSignUp(user *model.User) error {
	var hashedPw []byte
	var id uuid.UUID
	var err error

	utils.Info("%v", user)

	if hashedPw, err = bcrypt.GenerateFromPassword([]byte(user.Password), config.App.Controller.HashCost); err != nil {
		return err
	}

	if id, err = uuid.NewUUID(); err != nil {
		return err
	}

	user.Password = string(hashedPw)
	user.ID = id.String()

	err = c.base.store.User().Create(user)
	return err
}
