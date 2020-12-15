package domain

import (
	"reblog-server/dependency"
	"reblog-server/domain/interactors"
)

type interactor struct {
	user interactors.UserInteractor
	todo interactors.TodoInteractor
}

// NewInteractor ...
func NewInteractor() dependency.Interactor {
	return &interactor{
		user: interactors.UserInteractor{},
		todo: interactors.TodoInteractor{},
	}
}
