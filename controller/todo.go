package controller

import (
	"fmt"
	"reblog-server/model"
	"reblog-server/utils"

	"github.com/google/uuid"
)

type todoController struct {
	base *baseController
}

func newTodoController(base *baseController) *todoController {
	return &todoController{
		base: base,
	}
}

func (c *todoController) Get(id string) (*model.Todo, error) {
	var err error
	var todo *model.Todo

	todo, err = c.base.store.Todo().GetByID(id)
	if err != nil {
		utils.Error("[todoController] Failed to get Todo by id %s. Error: %s", id, err)
		return nil, err
	}

	return todo, nil
}

func (c *todoController) GetAll() ([]model.Todo, error) {
	var err error
	var todos []model.Todo

	todos, err = c.base.store.Todo().GetAll()
	if err != nil {
		utils.Error("[todoController] Failed to get all Todo. Error: %s", err)
		return nil, err
	}

	return todos, nil
}

func (c *todoController) Create(todo *model.Todo) (*Response, error) {
	var err error
	var id uuid.UUID

	if id, err = uuid.NewUUID(); err != nil {
		utils.Error("[todoController] Failed to generate UUID. Error: %s", err)
		return nil, err
	}

	idStrVal := id.String()
	todo.UUID = idStrVal
	res, err := c.base.store.Todo().Create(todo)
	if err != nil {
		utils.Error("[todoController] Failed to create todo. Error: %s", err)
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.Error("[todoController] Cannot get rows affected")
		return nil, err
	}

	return &Response{
		Message:      fmt.Sprintf("Created new Todo"),
		RowsAffected: int(rowsAffected),
	}, nil
}

func (c *todoController) Update(todo *model.Todo) (*Response, error) {
	var err error

	// Validate fields
	id := todo.ID
	if id == "" {
		utils.Error("[todoController] Can't update Todo due to missing ID field")
		return nil, &Error{
			Message: "Missing id field",
		}
	}

	res, err := c.base.store.Todo().UpdateByID(id, todo)
	if err != nil {
		utils.Error("[todoController] Failed to update todo")
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.Error("[todoController] Cannot get rows affected")
		return nil, err
	}

	return &Response{
		Message:      fmt.Sprintf("Updated Todo item of Id %s", id),
		RowsAffected: int(rowsAffected),
	}, nil
}

func (c *todoController) Delete(id string) (*Response, error) {
	var err error

	res, err := c.base.store.Todo().DeleteByID(id)
	if err != nil {
		utils.Error("[todoController] Failed to delete todo")
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.Error("[todoController] Cannot get rows affected")
		return nil, err
	}

	return &Response{
		Message:      fmt.Sprintf("Deleted Todo item of id %s", id),
		RowsAffected: int(rowsAffected),
	}, nil
}
