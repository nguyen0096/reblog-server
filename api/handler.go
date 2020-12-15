package api

import (
	"net/http"
)

var (
	baseHandler = BaseHandler{}
)

type BaseHandler struct {
}

func setInteractor() {
}

// Handler ...
type Handler struct {
	BaseHandler
	H func(w http.ResponseWriter, r *http.Request) (int, error)
}

func (ah Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if status, err := ah.H(w, r); err != nil {
		switch status {
		case http.StatusNotFound:
			http.Error(w, "You found nothing? Maybe another time", http.StatusInternalServerError)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func APIHandler(fn func(w http.ResponseWriter, r *http.Request) (int, error)) http.Handler {
	return &Handler{
		BaseHandler: baseHandler,
		H:           fn,
	}
}
