package apiv2

import (
	"net/http"
	"reblog-server/domain/interactor"
)

var (
	baseHandler = BaseHandler{}
)

type BaseHandler struct {
	Interactor interactor.Interactor
}

func setInteractor(iter interactor.Interactor) {
	baseHandler.Interactor = iter
}

// Handler ...
type Handler struct {
	BaseHandler
	H func(iter interactor.Interactor, w http.ResponseWriter, r *http.Request) (int, error)
}

func (ah Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if status, err := ah.H(ah.BaseHandler.Interactor, w, r); err != nil {
		switch status {
		case http.StatusNotFound:
			http.Error(w, "You found nothing? Maybe another time", http.StatusInternalServerError)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func APIHandler(fn func(iter interactor.Interactor, w http.ResponseWriter, r *http.Request) (int, error)) http.Handler {
	return &Handler{
		BaseHandler: baseHandler,
		H:           fn,
	}
}
