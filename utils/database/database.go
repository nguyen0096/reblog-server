package database

import (
	"fmt"
	"time"

	"reblog-server/utils"
	"reblog-server/utils/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewPostgresSQLConnection ...
func InitPostgres() *sqlx.DB {
	App := config.App

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		App.Database.Host,
		App.Database.Port,
		App.Database.User,
		App.Database.Password,
		App.Database.Name,
	)

	db, err := sqlx.Connect("postgres", connString)
	db.SetConnMaxIdleTime(time.Duration(App.Database.MaxIdleTimeConnection * float64(time.Second)))
	db.SetConnMaxLifetime(time.Duration(App.Database.MaxLifeTimeConnection * float64(time.Second)))
	db.SetMaxOpenConns(App.Database.MaxConnection)
	db.SetMaxIdleConns(App.Database.MinConnection)
	if err != nil {
		utils.Error("cannot initialize postgres connection. error: %s", err)
	}

	return db
}
