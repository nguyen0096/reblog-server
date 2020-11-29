package handler

import (
	"net/http"

	"reblog-server/domain/repository"
)

type HttpHandler struct {
	dbRepo repository.UserRepository
}

func (c *HttpHandler) InitHttpHandler() {
	http.Handle("/", http.FileServer(http.Dir("client")))
}

func handleLogin()

func NewHttpHandler(dbRepo repository.UserRepository) (*HttpHandler, error) {
	return &HttpHandler{
		dbRepo: dbRepo,
	}, nil
}
