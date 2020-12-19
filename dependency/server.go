package dependency

import (
	"github.com/jmoiron/sqlx"
)

// IServer ...
type IServer interface {
	// GETTERS
	Config() IConfig
	Database() *sqlx.DB

	// SETTERS
	SetDatabaseConnection(db *sqlx.DB)

	// methods
	Start()
}
