package handler

type Service interface {
	User() UserService
}
type UserService interface {
}
