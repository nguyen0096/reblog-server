package sqlstore

import (
	"reblog-server/dependency"

	"github.com/jmoiron/sqlx"
)

type sqlstore struct {
	db     *sqlx.DB
	stores stores
}

type stores struct {
	user *userSqlStore
}

func New(db *sqlx.DB) dependency.IStore {
	store := &sqlstore{
		db: db,
	}

	store.stores.user = NewUserSqlStore(store)

	return store
}

func (c *sqlstore) User() dependency.IUserStore {
	return c.stores.user
}
