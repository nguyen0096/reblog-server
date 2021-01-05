package sqlstore

import (
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

func (s userSqlStore) Get(username string) (*model.User, error) {
	user := model.User{}

	queryString := `SELECT id, username, first_name, last_name FROM rb_core.user WHERE user`
	row := s.base.db.QueryRow(queryString, username)
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
