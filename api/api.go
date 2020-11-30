package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes ...
type Routes struct {
	Root *mux.Router

	Users *mux.Router
}

// API ...
type API struct {
	BaseRoutes *Routes
}

// Init ...
func Init(root *mux.Router) *API {
	log.Println("init api")
	api := &API{
		BaseRoutes: &Routes{},
	}

	api.BaseRoutes.Root = root

	api.BaseRoutes.Users = root.PathPrefix("/users").Subrouter()

	api.InitUser()

	root.Handle("/{anything:.*}", http.HandlerFunc(api.Handle404))

	return api
}

// Handle404 ...
func (api *API) Handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
