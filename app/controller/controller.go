package controller

import (
	"reblog-server/app"
	"reblog-server/store"
)

type baseController struct {
	controllers

	store store.Store
}

type controllers struct {
	user *userController
}

// New ...
func New(store store.Store) app.Controller {
	base := baseController{}

	base.user = newUserController(base)

	return base
}

func (c baseController) User() app.UserController {
	return c.user
}
