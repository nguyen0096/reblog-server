package store

import (
	"reblog-server/model"
)

type Store interface {
	User() UserStore
}

type UserStore interface {
	Get(username string) (*model.User, error)
	Create(newUser model.User) error
}
