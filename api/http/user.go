package http

import (
	"net/http"

	"reblog-server/domain/model"
	"reblog-server/utils"
)

func (c *APIServer) initUserAPI() {
	c.User = c.Root.PathPrefix("/user").Subrouter()

	c.User.HandleFunc("", c.getUser).Methods("GET")
	c.User.HandleFunc("", c.createUser).Methods("POST")
}

func (api *APIServer) getUser(w http.ResponseWriter, r *http.Request) {
	utils.Info("GetUser hit!")
}

func (api *APIServer) createUser(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	var err error

	if user, err = model.UserFromJson(r.Body); err != nil {
		api.error(w, http.StatusBadRequest, err)
		return
	}

	if user == nil {
		api.error(w, http.StatusBadRequest, &APIError{Message: "Can't parse JSON data"})
		return
	}

	err = api.Service.User().CreateUserFromSignUp(user)
	if err != nil {
		api.error(w, http.StatusBadRequest, err)
		return
	}

	api.respond(w, http.StatusCreated, nil)
}
