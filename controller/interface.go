package controller

import (
	"reblog-server/model"
)

type Controller interface {
	User() IUserController
	Todo() ITodoController
}

type IUserController interface {
	CreateUserFromSignUp(user *model.User) error
	CreateToken(user *model.User) (string, error)
}

type ITodoController interface {
	Create(todo *model.Todo) (*Response, error)
	GetAll() ([]model.Todo, error)
	Update(todo *model.Todo) (*Response, error)
	Delete(id string) (*Response, error)
}
