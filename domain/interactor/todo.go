package interactor

import (
	"log"
	"reblog-server/domain/entity"
)

type TodoInteractor interface {
	GetAllTodos() ([]entity.Todo, error)
}

type todoInteractor struct {
}

func newTodoInteractor() TodoInteractor {
	return &todoInteractor{}
}

func (this *todoInteractor) GetAllTodos() ([]entity.Todo, error) {
	log.Println("interactor: GetAllTodos")
	return nil, nil
}
