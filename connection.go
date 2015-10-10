package main

import (
	"database/sql"
	"github.com/ranal/tarkvaraprojekt/Godeps/_workspace/src/github.com/lib/pq"
	"os"
)

func openDb() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("piletimyyk", connection)
	if err != nil {
		log.Println(err)
	}

	return db
}
