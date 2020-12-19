package sqlstore

import (
	"database/sql"
	"fmt"
	"reblog-server/domain/entity"
)

type userSqlStore struct {
	store *sqlstore
}

func NewUserSqlStore(store *sqlstore) *userSqlStore {
	return &userSqlStore{
		store: store,
	}
}

func (s *userSqlStore) GetUserById(id string) (*entity.User, error) {
	user := entity.User{}

	queryString := `SELECT id, username, first_name, last_name FROM rb_core.user WHERE id=$1`
	row := s.store.db.QueryRow(queryString, id)
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
