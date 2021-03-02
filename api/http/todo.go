package http

import (
	"net/http"
	"reblog-server/domain/model"

	"github.com/gorilla/mux"
)

func (c *APIServer) initTodoAPI() {
	c.Todo = c.Root.PathPrefix("/todo").Subrouter()

	c.Todo.HandleFunc("/{id}", c.getTodo).Methods("GET")
	c.Todo.HandleFunc("", c.getAllTodo).Methods("GET")
	c.Todo.HandleFunc("", c.createTodo).Methods("POST")
	c.Todo.HandleFunc("", c.updateTodo).Methods("PUT")
	c.Todo.HandleFunc("/{id}", c.deleteTodo).Methods("DELETE")
}

func (api *APIServer) getTodo(w http.ResponseWriter, r *http.Request) {

}

func (api *APIServer) getAllTodo(w http.ResponseWriter, r *http.Request) {
	var err error
	var todos []model.Todo

	todos, err = api.Service.Todo().GetAll()
	if err != nil {
		api.error(w, http.StatusInternalServerError, err)
		return
	}

	api.respond(w, http.StatusOK, todos)
}

func (api *APIServer) createTodo(w http.ResponseWriter, r *http.Request) {
	var err error
	var todo *model.Todo

	todo, err = model.TodoFromJSON(r.Body)
	if err != nil {
		api.error(w, http.StatusBadRequest, err)
		return
	}

	res, err := api.Service.Todo().Create(todo)
	if err != nil {
		api.error(w, http.StatusInternalServerError, err)
		return
	}

	api.respond(w, http.StatusOK, res)
}

func (api *APIServer) updateTodo(w http.ResponseWriter, r *http.Request) {
	var err error
	var todo *model.Todo

	todo, err = model.TodoFromJSON(r.Body)
	if err != nil {
		api.error(w, http.StatusBadRequest, err)
		return
	}

	res, err := api.Service.Todo().Update(todo)
	if err != nil {
		api.error(w, http.StatusInternalServerError, err)
		return
	}

	api.respond(w, http.StatusOK, res)
}

func (api *APIServer) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	res, err := api.Service.Todo().Delete(id)
	if err != nil {
		api.error(w, http.StatusInternalServerError, err)
		return
	}

	api.respond(w, http.StatusOK, res)
}
