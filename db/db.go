package db

import (
	"database/sql"
	"os"
)

// GetConnection Create a connection with database
func GetConnection() *sql.DB {
	connStr := os.Getenv("POSTGRESQL_URL")
	dbase, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}

	return dbase
}
