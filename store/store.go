package store

import (
	"reblog-server/model"
)

type Store interface {
	User() UserStore
}

type UserStore interface {
	Get(id string) (*model.User, error)
}
