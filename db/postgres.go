package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	dbUsername := os.Getenv("POSTGRESUSERNAME")
	dbPassword := os.Getenv("POSTGRESPASSWORD")
	dbName := os.Getenv("DATABASENAME")

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUsername, dbPassword, dbName))
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database: ", err)
	}

	return db, nil
}
