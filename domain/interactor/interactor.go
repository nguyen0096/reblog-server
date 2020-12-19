package interactor

import (
	"reblog-server/dependency"
)

type baseInteractor struct {
	store       dependency.IStore
	interactors interactors
}

type interactors struct {
	todo TodoInteractor
	user UserInteractor
}

// New ...
func New(store dependency.IStore) dependency.IInteractor {
	base := &baseInteractor{
		store: store,
	}

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
