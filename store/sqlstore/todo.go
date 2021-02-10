package sqlstore

import (
	"database/sql"
	"fmt"
	"reblog-server/model"
	"reblog-server/store/orm"
	"reblog-server/utils"
	"reflect"
	"strings"
)

type todoSqlStore struct {
	base *baseSqlStore
}

func newTodoStore(store *baseSqlStore) *todoSqlStore {
	return &todoSqlStore{
		base: store,
	}
}

func (c todoSqlStore) Create(newTodo *model.Todo) (sql.Result, error) {
	var err error

	utils.Info("%v", newTodo)

	queryStr := fmt.Sprintf("INSERT INTO todo (uuid, title, short_description, description, created_by) VALUES ($1, $2, $3, $4, $5)")

	res, err := c.base.db.Exec(queryStr, newTodo.UUID, newTodo.Title, newTodo.ShortDescription.String, newTodo.Description.String, newTodo.CreatedBy.String)
	if err != nil {
		utils.Error("[todoSqlStore] Failed to exec query")
		return nil, err
	}

	return res, nil
}

func (c todoSqlStore) GetByID(id string) (*model.Todo, error) {
	var err error
	var todo *model.Todo

	row := c.base.db.QueryRowx("SELECT * FROM todo WHERE id=$1", id)
	err = row.StructScan(&todo)
	if err != nil {
		utils.Error("[todoSqlStore] Failed to scan struct")
		return nil, err
	}

	return todo, nil
}

func (c todoSqlStore) GetAll() ([]model.Todo, error) {
	var err error
	var todos []model.Todo

	rows, err := c.base.db.Queryx("SELECT * FROM todo")
	if err != nil {
		utils.Error("[todoSqlStore] Failed to query")
	}

	for rows.Next() {
		todo := model.Todo{}
		err = rows.StructScan(&todo)
		if err != nil {
			utils.Error("[todoSqlStore] Failed to scan struct")
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (c todoSqlStore) UpdateByID(id string, todo *model.Todo) (sql.Result, error) {
	var err error
	// var query string = "UPDATE todo SET "
	var queryParts []string
	queryParts = append(queryParts, "UPDATE todo SET")

	t := reflect.TypeOf(*todo)
	v := reflect.ValueOf(*todo)
	for i := 1; i < t.NumField(); i++ {
		fVal, ok := v.Field(i).Interface().(orm.DataType)
		if !ok {
			utils.Info("Not FieldType. %s", t.Field(i).Name)
			continue
		}

		if fVal.IsNull() {
			continue
		}

		fType := t.Field(i)
		tag := fType.Tag.Get("db")
		if tag == "id" {
			// Remove id from updated fields
			continue
		}

		queryParts = append(queryParts, fmt.Sprintf("%s='%s'", tag, fVal.Value()))
	}
	queryParts = append(queryParts, fmt.Sprintf("WHERE id=%s", id))
	queryStr := strings.Join(queryParts, " ")

	utils.Info("%s", queryStr)

	res, err := c.base.db.Exec(queryStr)
	if err != nil {
		utils.Error("[todoSqlStore] Failed to exec query: %s", queryStr)
		return nil, err
	}

	return res, nil
}

func (c todoSqlStore) DeleteByID(id string) (sql.Result, error) {
	var err error

	res, err := c.base.db.Exec("DELETE FROM todo WHERE id=$1", id)
	if err != nil {
		utils.Error("[todoSqlStore] Failed to exec query")
		return nil, err
	}

	return res, nil
}
