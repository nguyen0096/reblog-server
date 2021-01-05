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

func (s userSqlStore) Create(newUser model.User) error {
	queryString := `INSERT INTO rb_core.user (username, passwor) VALUES (?, ?)`
	_, err := s.base.db.Exec(queryString, newUser.Username, newUser.Password)
	if err != nil {
		return err
	}
	return nil
}
