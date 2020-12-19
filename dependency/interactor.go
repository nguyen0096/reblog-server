package dependency

// Interactor ...
type Interactor interface {
	User() IUserInteractor
	Todo() ITodoInteractor
}

type IUserInteractor interface {
}

type ITodoInteractor interface {
}
