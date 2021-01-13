package sqlstore

import (
	"log"
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

func (s userSqlStore) GetByUsername(username string) (*model.User, error) {
	user := model.User{}

	// runtime.Breakpoint()

	// TODO: scan empty column
	queryString := `SELECT id, username, password FROM rb_core.user WHERE username=$1`
	row := s.base.db.QueryRowx(queryString, username)
	err := row.StructScan(&user)

	if err != nil {
		log.Printf("failed to scan row. err: %v", err)
		return nil, err
	}

	return &user, nil
}

func (s userSqlStore) Create(newUser *model.User) error {
	queryString := `INSERT INTO rb_core.user (id, username, password) VALUES ($1, $2, $3);`
	_, err := s.base.db.Exec(queryString, newUser.ID, newUser.Username, newUser.Password)
	if err != nil {
		return err
	}

	// log.Printf("%T", res)

	return nil
}
