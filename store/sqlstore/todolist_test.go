package sqlstore

import (
	"database/sql"
	"testing"

	"reblog-server/domain/model"
	"reblog-server/utils/config"
	"reblog-server/utils/database"
)

func TestTodoListStore(t *testing.T) {
	config.InitConfig()

	gormConn := database.NewPostgresGormConn()

	s := newTodoListStore(gormConn)

	newList := &model.TodoList{
		Name:        "Test todolist 1",
		Description: sql.NullString{String: "Test todolist 1 - desc", Valid: true},
	}

	s.Create(newList)
}
