package dependency

import (
	"reblog-server/model"
)

type IStore interface {
	User() IUserStore
}

type IUserStore interface {
	GetUserById(id string) (*model.User, error)
}
