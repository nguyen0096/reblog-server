package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Error      error  `json:"-"`
}

func (c *APIServer) debug(format string, args ...interface{}) {
	// TODO implement debug handler for API
}

func (c *APIServer) respond(writer http.ResponseWriter, code int, src interface{}) {
	var body []byte
	var err error

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if body, err = json.Marshal(src); err != nil {
		errorBody := "{\"status\": 500, \"message\": \"Something happened wrong during generating response\"}"
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(errorBody))
		return
	}

	writer.WriteHeader(code)
	writer.Write(body)
}

// func (c *APIServer) error(w http.ResponseWriter, code int, message string) {

// }

func (c *APIServer) error(w http.ResponseWriter, err error) {
	var statusCode int
	var message string

	switch e := err.(type) {
	case *json.UnsupportedTypeError, *json.UnmarshalTypeError, *json.SyntaxError:
		statusCode = http.StatusBadRequest
		message = "Request body is invalid"
	default:
		statusCode = http.StatusInternalServerError
		message = e.Error()
	}

	errRes := &errorResponse{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}

	log.Printf("%v", errRes)

	c.respond(w, statusCode, errRes)
}
