package controller

import (
	"log"
	"reblog-server/store"
)

type baseController struct {
	// Children
	controllers

	// Dependencies
	store store.Store
}

type controllers struct {
	user *userController
	todo *todoController
}

// New ...
func New(store store.Store) Controller {
	if store == nil {
		log.Panicf("nil store param")
	}

	base := &baseController{
		store: store,
	}

	base.user = newUserController(base)
	base.todo = newTodoController(base)

	return base
}

// Implement interface
func (c baseController) User() IUserController {
	return c.user
}

func (c baseController) Todo() ITodoController {
	return c.todo
}
