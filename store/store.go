package store

import (
	"reblog-server/model"
)

type Store interface {
	User() UserStore
}

type UserStore interface {
	GetByUsername(username string) (*model.User, error)
	Create(newUser *model.User) error
}
