package infra

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"reblog-server/dependency"
)

// NewPostgresConnection ...
func NewPostgresConnection(conf dependency.IConfig) *sqlx.DB {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.GetHostname(),
		conf.GetPort(),
		conf.GetUser(),
		conf.GetPassword(),
		conf.GetDatabase())

	db, err := sqlx.Connect("postgres", connString)
	// db.SetConnMaxIdleTime(time.Duration(databaseConfig.MaxIdleTimeConnection * float64(time.Second)))
	// db.SetConnMaxLifetime(time.Duration(databaseConfig.MaxLifeTimeConnection * float64(time.Second)))
	// db.SetMaxOpenConns(databaseConfig.MaxConnection)
	// db.SetMaxIdleConns(databaseConfig.MinConnection)
	if err != nil {
		log.Fatalf("Failed to init database connection. [error] %v", err)
	}

	return db
}
