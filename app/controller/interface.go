package controller

import (
	"reblog-server/model"
)

type IController interface {
	User() IUserController
}

type IUserController interface {
	CreateUserFromSignUp(user *model.User) error
	CreateToken(user *model.User) (string, error)
}
