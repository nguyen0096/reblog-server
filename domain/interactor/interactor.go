package interactor

import (
	"reblog-server/dependency"
)

type baseInteractor struct {
	interactors interactors
}

type interactors struct {
	todo TodoInteractor
	user UserInteractor
}

// NewInteractor ...
func NewInteractor() dependency.Interactor {
	base := &baseInteractor{}

	base.interactors.user = UserInteractor{}
	base.interactors.todo = TodoInteractor{}

	return base
}

func (c *baseInteractor) Todo() dependency.ITodoInteractor {
	return c.interactors.todo
}

func (c *baseInteractor) User() dependency.IUserInteractor {
	return c.interactors.user
}
