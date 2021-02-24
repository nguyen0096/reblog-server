package todo

import (
	"fmt"
	"reblog-server/domain/model"
	"reblog-server/store"
	"reblog-server/utils"

	"github.com/google/uuid"
)

type TodoService interface {
	Create(todo *model.Todo) (*utils.Response, error)
	GetAll() ([]model.Todo, error)
	Update(todo *model.Todo) (*utils.Response, error)
	Delete(id string) (*utils.Response, error)
}

type todoService struct {
	store store.TodoStore
}

func NewTodoService(store store.TodoStore) TodoService {
	return &todoService{
		store: store,
	}
}

func (c *todoService) Get(id string) (*model.Todo, error) {
	var err error
	var todo *model.Todo

	todo, err = c.store.GetByID(id)
	if err != nil {
		utils.Error("[todoController] Failed to get Todo by id %s. Error: %s", id, err)
		return nil, err
	}

	return todo, nil
}

func (c *todoService) GetAll() ([]model.Todo, error) {
	var err error
	var todos []model.Todo

	todos, err = c.store.GetAll()
	if err != nil {
		utils.Error("[todoController] Failed to get all Todo. Error: %s", err)
		return nil, err
	}

	return todos, nil
}

func (c *todoService) Create(todo *model.Todo) (*utils.Response, error) {
	var err error
	var id uuid.UUID

	if id, err = uuid.NewUUID(); err != nil {
		utils.Error("[todoController] Failed to generate UUID. Error: %s", err)
		return nil, err
	}

	idStrVal := id.String()
	todo.UUID = idStrVal
	res, err := c.store.Create(todo)
	if err != nil {
		utils.Error("[todoController] Failed to create todo. Error: %s", err)
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.Error("[todoController] Cannot get rows affected")
		return nil, err
	}

	return &utils.Response{
		Message:      fmt.Sprintf("Created new Todo"),
		RowsAffected: int(rowsAffected),
	}, nil
}

func (c *todoService) Update(todo *model.Todo) (*utils.Response, error) {
	var err error

	// Validate fields
	id := todo.ID
	if id == "" {
		utils.Error("[todoController] Can't update Todo due to missing ID field")
		return nil, &utils.AppError{
			Message: "Missing id field",
		}
	}

	res, err := c.store.UpdateByID(id, todo)
	if err != nil {
		utils.Error("[todoController] Failed to update todo")
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.Error("[todoController] Cannot get rows affected")
		return nil, err
	}

	return &utils.Response{
		Message:      fmt.Sprintf("Updated Todo item of Id %s", id),
		RowsAffected: int(rowsAffected),
	}, nil
}

func (c *todoService) Delete(id string) (*utils.Response, error) {
	var err error

	res, err := c.store.DeleteByID(id)
	if err != nil {
		utils.Error("[todoController] Failed to delete todo")
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.Error("[todoController] Cannot get rows affected")
		return nil, err
	}

	return &utils.Response{
		Message:      fmt.Sprintf("Deleted Todo item of id %s", id),
		RowsAffected: int(rowsAffected),
	}, nil
}
