package controller

import "reblog-server/dto"

type IController interface {
	User() IUserController
}

type IUserController interface {
	CreateUserFromSignUp(dto.LoginForm) error
}
