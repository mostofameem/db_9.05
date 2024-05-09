package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB
var err error

func InitDB() error {
	Db, err = sqlx.Connect("postgres", "postgresql://root:admin@localhost:5432/root?sslmode=disable")
	if err != nil {
		return err
	}

	// Ping the database to check if the connection is successful
	if err := Db.Ping(); err != nil {
		return err
	}

	return nil
}
