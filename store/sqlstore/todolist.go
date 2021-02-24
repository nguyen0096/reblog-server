// TodoListStore uses gorm instead of using sqlx directly

package sqlstore

import (
	"database/sql"
	"reblog-server/domain/model"
	"reblog-server/utils/database"

	"gorm.io/gorm"
)

type todoListStore struct {
	db *gorm.DB
}

func newTodoListStore(db *gorm.DB) *todoListStore {
	return &todoListStore{
		db: db,
	}
}

func (c todoListStore) MigrateModels() {

}

func (c todoListStore) Create(newList *model.TodoList) (sql.Result, error) {
	result := c.db.Create(newList)

	return &database.SqlResult{
		Error:           result.Error,
		RowsAffectedVal: result.RowsAffected,
	}, nil
}
