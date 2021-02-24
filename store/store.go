package store

import (
	"database/sql"
	"reblog-server/domain/model"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Store interface {
	SetGormConn(conn *gorm.DB)
	SetSqlxConn(conn *sqlx.DB)

	Migrate()

	// Get stores
	User() UserStore
	Todo() TodoStore
	TodoList() TodoListStore
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

type TodoListStore interface {
	Create(newList *model.TodoList) (sql.Result, error)
}
