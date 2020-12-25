package sqlstore

import (
	"database/sql"
	"fmt"
	"reblog-server/model"
)

type userSqlStore struct {
	base *baseSqlStore
}

func newUserStore(store *baseSqlStore) *userSqlStore {
	return &userSqlStore{
		base: store,
	}
}

func (s userSqlStore) Get(id string) (*model.User, error) {
	user := model.User{}

	queryString := `SELECT id, username, first_name, last_name FROM rb_core.user WHERE id=$1`
	row := s.base.db.QueryRow(queryString, id)
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	return &user, nil
}
