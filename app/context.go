package app

import (
	"reblog-server/domain/interactor"
)

type appContext struct {
	Interactor interactor.Interactor
}
