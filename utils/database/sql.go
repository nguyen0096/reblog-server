package database

import (
	"fmt"
	"time"

	"reblog-server/utils"
	"reblog-server/utils/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlResult struct {
	Error           error
	RowsAffectedVal int64
	LastInsertIdVal int64
}

func (c *SqlResult) RowsAffected() (int64, error) {
	return c.RowsAffectedVal, c.Error
}

func (c *SqlResult) LastInsertId() (int64, error) {
	return c.LastInsertIdVal, c.Error
}

func dbDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.App.Database.Host,
		config.App.Database.Port,
		config.App.Database.User,
		config.App.Database.Password,
		config.App.Database.Name,
	)
}

// NewPostgresSqlxConn ...
func NewPostgresSqlxConn() *sqlx.DB {
	db, err := sqlx.Connect("postgres", dbDSN())
	db.SetConnMaxIdleTime(time.Duration(config.App.Database.MaxIdleTimeConnection * float64(time.Second)))
	db.SetConnMaxLifetime(time.Duration(config.App.Database.MaxLifeTimeConnection * float64(time.Second)))
	db.SetMaxOpenConns(config.App.Database.MaxConnection)
	db.SetMaxIdleConns(config.App.Database.MinConnection)
	if err != nil {
		utils.Error("cannot initialize postgres connection. error: %s", err)
	}

	return db
}

// NewPostgresGormConn return instance of gorm.DB connection
func NewPostgresGormConn() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbDSN()), &gorm.Config{})
	if err != nil {
		utils.Error("cannot initialize postgres gorm connection. error: %s", err)
	}
	return db
}
