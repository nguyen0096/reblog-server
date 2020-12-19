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
	SetStore(store IStore)
	SetInteractor(iter IInteractor)

	// methods
	Start()
}
