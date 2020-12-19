package dependency

// IInteractor ...
type IInteractor interface {
	User() IUserInteractor
	Todo() ITodoInteractor
}

type IUserInteractor interface {
}

type ITodoInteractor interface {
}
