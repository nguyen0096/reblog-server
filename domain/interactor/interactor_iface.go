package interactor

// Interactor ...
type Interactor interface {
	User() UserInteractor
	Todos() TodoInteractor
}

type interactor struct {
	user UserInteractor
	todo TodoInteractor
}

// NewInteractor ...
func NewInteractor() Interactor {
	return &interactor{
		user: newUserInteractor(),
		todo: newTodoInteractor(),
	}
}

func (iter *interactor) User() UserInteractor {
	return iter.user
}

func (iter *interactor) Todos() TodoInteractor {
	return iter.todo
}
