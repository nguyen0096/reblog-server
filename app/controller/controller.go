package controller

import "reblog-server/app"

type baseController struct {
	controllers
}

type controllers struct {
	user *userController
}

// New ...
func New() app.Controller {
	base := baseController{}

	base.user = newUserController(base)

	return base
}

func (c baseController) User() app.UserController {
	return c.user
}
