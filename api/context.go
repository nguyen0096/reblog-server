package api

import (
	"net/http"
	"reblog-server/dependency"
)

type Context struct {
	Server     dependency.IServer
	Request    *http.Request
	Writer     http.ResponseWriter
	TestString string
}
