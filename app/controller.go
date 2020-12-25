package app

type Controller interface {
	User() UserController
}

type UserController interface {
}
