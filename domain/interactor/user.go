package interactor

import (
	"reblog-server/domain/entity"
)

// UserInteractor ...
type UserInteractor interface {
	GetAllUsers() ([]entity.User, error)
}

type userInteractor struct {
}

func newUserInteractor() UserInteractor {
	return &userInteractor{}
}

func (iter *userInteractor) GetAllUsers() ([]entity.User, error) {
	return nil, nil
}
