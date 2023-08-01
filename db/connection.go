package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() {
	var err error
	connStr := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
}

func CloseDatabase() {
	if db != nil {
		db.Close()
	}
}
