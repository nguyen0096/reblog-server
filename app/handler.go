package app

import (
	"net/http"
)

// MyHandler ...
type MyHandler struct {
	Handler func(w http.ResponseWriter, r *http.Request) error
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.Handler(w, r); err != nil {

	}
}
