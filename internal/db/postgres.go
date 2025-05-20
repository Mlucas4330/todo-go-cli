package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "user=postgres dbname=tasks sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	return db
}
