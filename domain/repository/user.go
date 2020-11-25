package repository

import (
	"reblog-auth/domain/entity"
)

type UserRepository interface {
	GetUserByID(id string) (*entity.User, error)
}
