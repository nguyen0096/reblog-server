package controller

import (
	"reblog-server/app"
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
}

// New ...
func New(store store.Store) app.Controller {
	base := baseController{}

	base.user = newUserController(base)

	return base
}

// Implement interface
func (c baseController) User() app.UserController {
	return c.user
}
