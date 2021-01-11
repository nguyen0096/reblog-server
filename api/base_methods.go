package api

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Error []error `json:"error"`
}

func (c *APIServer) error(w http.ResponseWriter, code int, e error) {
	var body []byte
	var err error

	res := errorResponse{
		Error: []error{e},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if body, err = json.Marshal(res); err != nil {
		errorBody := "{\"status\": 500, \"message\": \"Something happened wrong during generating response\"}"
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errorBody))
		return
	}

	w.WriteHeader(code)
	w.Write(body)
}

func (c *APIServer) respond(w http.ResponseWriter, code int, data interface{}) {
	var body []byte
	var err error

	res := response{
		Data: data,
	}

	if body, err = json.Marshal(res); err != nil {
		errorBody := "{\"status\": 500, \"message\": \"Something happened wrong during generating response\"}"
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errorBody))
		return
	}

	w.WriteHeader(code)
	w.Write(body)
}

func (c *APIServer) debug(format string, args ...interface{}) {
	// TODO implement debug handler for API
}
