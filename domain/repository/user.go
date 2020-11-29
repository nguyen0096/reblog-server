package repository

import (
	"reblog-server/domain/entity"
)

type UserRepository interface {
	GetUserByID(id string) (*entity.User, error)
}
