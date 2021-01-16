package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewPostgresSQLConnection ...
func NewPostgresSQLConnection() *sqlx.DB {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		App.Database.Host,
		App.Database.Port,
		App.Database.User,
		App.Database.Password,
		App.Database.Name,
	)

	db, err := sqlx.Connect("postgres", connString)
	// db.SetConnMaxIdleTime(time.Duration(databaseConfig.MaxIdleTimeConnection * float64(time.Second)))
	// db.SetConnMaxLifetime(time.Duration(databaseConfig.MaxLifeTimeConnection * float64(time.Second)))
	// db.SetMaxOpenConns(databaseConfig.MaxConnection)
	// db.SetMaxIdleConns(databaseConfig.MinConnection)
	if err != nil {
		log.Fatalf("error initialize postgres connection. error: %v", err)
	}

	return db
}
