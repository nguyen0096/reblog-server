package model

import (
	"encoding/json"
	"io"

	"reblog-server/store/orm"
)

// Todo uses sqlstore
type Todo struct {
	UUID             string         `json:"-" db:"uuid"`
	ID               string         `json:"id,omitempty" db:"id"`
	Title            string         `json:"title,omitempty" db:"title"`
	ShortDescription orm.NullString `json:"short_description,omitempty" db:"short_description"`
	Description      orm.NullString `json:"description,omitempty" db:"description"`
	CreatedBy        orm.NullString `json:"created_by,omitempty" db:"created_by"`
}

func TodoFromJSON(data io.Reader) (*Todo, error) {
	var todo *Todo
	err := json.NewDecoder(data).Decode(&todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
