package orm

import "time"

// Basic model
type Model struct {
	ID        uint
	CreateAt  time.Time
	UpdatedAt time.Time
}
