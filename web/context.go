package web

import (
	"reblog-server/domain/interactor"
)

// Context ...
type Context struct {
	Interactor interactor.Interactor
}
