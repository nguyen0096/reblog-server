package app

import "reblog-server/dto"

type Controller interface {
	User() UserController
}

type UserController interface {
	CreateUser(dto.LoginForm) error
}
