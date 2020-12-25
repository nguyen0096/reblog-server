package handler

import "reblog-server/dependency"

type service struct {
	user *userService
}

func NewService() dependency.IService {
	return service{
		user: NewUserService()
	}
}
