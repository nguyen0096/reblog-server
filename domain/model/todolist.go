package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type TodoList struct {
	gorm.Model
	Name        string
	Active      bool
	Description sql.NullString
	DueDate     time.Time
	// Owner -> TBD
}
