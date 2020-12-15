package dependency

// Interactor ...
type Interactor interface {
	IUserInteractor
	ITodoInteractor
}

type IUserInteractor interface {
}

type ITodoInteractor interface {
}
