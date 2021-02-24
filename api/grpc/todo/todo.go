package todo

import (
	"context"
	"reblog-server/domain/model"
	"reblog-server/service/todo"
	"reblog-server/store/orm"
	"reblog-server/utils"
)

type TodoHandler struct {
	Service todo.TodoService
}

func (c TodoHandler) AddTodo(ctx context.Context, t *TodoItem) (*AddTodoResponse, error) {
	utils.Info("gRPC endpoint hit!")
	todo := model.Todo{
		Title: t.Title,
		Description: orm.NullString{
			String: t.Description,
		},
		ShortDescription: orm.NullString{
			String: t.Description,
		},
		CreatedBy: orm.NullString{
			String: t.CreatedBy,
		},
	}

	res, err := c.Service.Create(&todo)
	if err != nil {
		return nil, err
	}

	return &AddTodoResponse{
		AddedTodo: int32(res.RowsAffected),
		Message:   res.Message,
	}, nil
}
