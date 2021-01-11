package controller

type AppError struct {
	Code    int
	Message string
	Detail  string
	Where   string
	Error   error
}
