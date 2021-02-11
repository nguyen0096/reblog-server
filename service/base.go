package service

import (
	"reblog-server/store"
	"reblog-server/utils"
)

type App interface {
	User() UserController
	Todo() TodoController
}

type baseService struct {
	store store.Store

	services
}

type services struct {
	user *userController
	todo *todoController
}

// New return a base controller
func New(store store.Store) App {
	if store == nil {
		utils.Panic("nil store param")
	}

	base := &baseService{
		store: store,
	}

	base.user = newUserController(base)
	base.todo = newTodoController(base)

	return base
}

// Interface implementation
func (c baseService) User() UserController {
	return c.user
}

func (c baseService) Todo() TodoController {
	return c.todo
}
