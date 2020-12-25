package api

import "github.com/gorilla/mux"

type Routes struct {
	Root *Router

	User *Router
}

type Router struct {
	Mux         *mux.Router
	Children    []Router
	Middlewares []HandlerFunc
}
