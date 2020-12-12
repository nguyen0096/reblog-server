package web

import (
	"net/http"
)

// Handler ...
type Handler struct {
	Context *Context
	H       func(ctx *Context, w http.ResponseWriter, r *http.Request) (int, error)
}

func (ah Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if status, err := ah.H(ah.Context, w, r); err != nil {
		switch status {
		case http.StatusNotFound:
			http.Error(w, "You found nothing? Maybe another time", http.StatusInternalServerError)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

// NewHandler ...
func NewHandler(fn func(c *Context, w http.ResponseWriter, r *http.Request) (int, error)) *Handler {

	ctx := &Context{}

	return &Handler{
		Context: ctx,
		H:       fn,
	}
}
