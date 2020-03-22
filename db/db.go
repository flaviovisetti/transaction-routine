package db

import (
	"os"

	"github.com/jinzhu/gorm"
)

// DBCon it is a pointer to gorm.DB
var DBCon *gorm.DB

// InitDB start connection with database
func InitDB() {
	var err error

	connStr := os.Getenv("POSTGRESQL_URL")
	DBCon, err = gorm.Open("postgres", connStr)

	if err != nil {
		panic("Failed to connect with database")
	}
}
