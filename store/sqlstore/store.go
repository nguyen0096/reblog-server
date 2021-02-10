package sqlstore

import (
	"reblog-server/store"

	"github.com/jmoiron/sqlx"
)

type baseSqlStore struct {
	db     *sqlx.DB
	stores *sqlStores
}

type sqlStores struct {
	user *userSqlStore
	todo *todoSqlStore
}

func New(db *sqlx.DB) store.Store {
	store := &baseSqlStore{
		db:     db,
		stores: &sqlStores{},
	}

	store.stores.user = newUserStore(store)
	store.stores.todo = newTodoStore(store)

	return store
}

func (c *baseSqlStore) User() store.UserStore {
	return c.stores.user
}

func (c *baseSqlStore) Todo() store.TodoStore {
	return c.stores.todo
}
