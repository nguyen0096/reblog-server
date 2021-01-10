package controller

import "reblog-server/dto"

type IController interface {
	User() IUserController
}

type IUserController interface {
	CreateUserFromSignup(dto.LoginForm) error
}
