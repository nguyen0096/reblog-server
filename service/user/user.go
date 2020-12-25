package user

import "reblog-server/dependency"

type userService struct{}

func NewUserService() dependency.IUserInteractor {
	return &userService{}
}

func (c *userService) CreateUser() {

}
