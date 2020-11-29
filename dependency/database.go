package dependency

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"reblog-server/config"
)

func NewPostgreSQLConnection() (*sqlx.DB, error) {
	databaseConfig := config.AppConfig.Database
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Name)
	fmt.Println(connString)
	db, err := sqlx.Connect("postgres", connString)
	// db.SetConnMaxIdleTime(time.Duration(databaseConfig.MaxIdleTimeConnection * float64(time.Second)))
	// db.SetConnMaxLifetime(time.Duration(databaseConfig.MaxLifeTimeConnection * float64(time.Second)))
	// db.SetMaxOpenConns(databaseConfig.MaxConnection)
	// db.SetMaxIdleConns(databaseConfig.MinConnection)
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}
