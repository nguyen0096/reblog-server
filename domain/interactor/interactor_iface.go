package interactor

// Interactor ...
type Interactor interface {
	User() UserInteractor
}

type interactor struct {
	user UserInteractor
}

// NewInteractor ...
func NewInteractor() Interactor {
	return &interactor{
		user: newUserInteractor(),
	}
}

func (iter *interactor) User() UserInteractor {
	return iter.user
}
