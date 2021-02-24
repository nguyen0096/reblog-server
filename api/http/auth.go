package http

import (
	"encoding/json"
	"net/http"

	"reblog-server/domain/model"
	"reblog-server/utils"
)

func (c *APIServer) initAuthAPI() {
	c.Auth = c.Root.PathPrefix("/auth").Subrouter()

	c.Auth.HandleFunc("/login", c.handleLogin).Methods("POST")
}

func (api *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	var u model.User
	var err error

	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		utils.Error("Failed to parse json", err)
		api.error(w, http.StatusBadRequest, err)
		return
	}

	token, err := api.Service.User().CreateToken(&u)
	if err != nil {
		api.error(w, http.StatusBadRequest, err)
		return
	}

	api.respond(w, http.StatusOK, token)
	return
}
