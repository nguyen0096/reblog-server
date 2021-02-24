package sqlstore

import (
	"reblog-server/domain/model"
	"reblog-server/store"
	"reblog-server/utils"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type baseSqlStore struct {
	gormConn *gorm.DB
	sqlxConn *sqlx.DB
	stores   *sqlStores
}

type sqlStores struct {
	user     *userSqlStore
	todo     *todoSqlStore
	todoList *todoListStore
}

// Implement Store interface
func (c *baseSqlStore) User() store.UserStore {
	return c.stores.user
}

func (c *baseSqlStore) Todo() store.TodoStore {
	return c.stores.todo
}

// TodoList return TodoListStore
// design-pattern: singleton
func (c *baseSqlStore) TodoList() store.TodoListStore {
	if c.stores.todoList == nil {
		if c.gormConn == nil {
			utils.Error("cannot init new TodoList store. gorm conn is nil")
		}
		c.stores.todoList = newTodoListStore(c.gormConn)
	}
	return c.stores.todoList
}

func (c *baseSqlStore) SetGormConn(conn *gorm.DB) {
	c.gormConn = conn
}

func (c *baseSqlStore) SetSqlxConn(conn *sqlx.DB) {
	c.sqlxConn = conn
}

// New SqlStore from given DB connection
func New() store.Store {
	store := &baseSqlStore{
		stores: &sqlStores{},
	}

	store.stores.user = newUserStore(store)
	store.stores.todo = newTodoStore(store)

	return store
}

func (c *baseSqlStore) Migrate() {
	if c.gormConn == nil {
		utils.Error("cannot migrate model. gorm connection is nil")
	}

	c.gormConn.AutoMigrate(&model.TodoList{})
}
