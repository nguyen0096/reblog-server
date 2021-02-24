package store

import (
	"database/sql"
	"reblog-server/domain/model"
)

type Store interface {
	User() UserStore
	Todo() TodoStore
}

type UserStore interface {
	GetByUsername(username string) (*model.User, error)
	Create(newUser *model.User) error
}

type TodoStore interface {
	// Generic CRUD
	Create(newTodo *model.Todo) (sql.Result, error)
	GetByID(id string) (*model.Todo, error)
	GetAll() ([]model.Todo, error)
	UpdateByID(id string, todo *model.Todo) (sql.Result, error)
	DeleteByID(id string) (sql.Result, error)
}
