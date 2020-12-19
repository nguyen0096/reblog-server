package dependency

import "reblog-server/domain/entity"

type IStore interface {
	User() IUserStore
}

type IUserStore interface {
	GetUserById(id string) (*entity.User, error)
}
