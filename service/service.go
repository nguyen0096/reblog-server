package service

import (
	"reblog-server/service/todo"
	"reblog-server/service/user"
	"reblog-server/store"
	"reblog-server/utils"
)

type App interface {
	User() user.UserService
	Todo() todo.TodoService
}

type appService struct {
	store store.Store

	services
}

type services struct {
	user user.UserService
	todo todo.TodoService
}

// New return a base controller
func New(store store.Store) App {
	if store == nil {
		utils.Panic("nil store param")
	}

	base := &appService{
		store: store,
	}

	base.user = user.NewUserService(store.User())
	base.todo = todo.NewTodoService(store.Todo())

	return base
}

func (c *appService) User() user.UserService {
	return c.user
}

func (c *appService) Todo() todo.TodoService {
	return c.todo
}
