package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func Connect () *sql.DB {
	connStr := "user=root password=secret dbname=koala sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
