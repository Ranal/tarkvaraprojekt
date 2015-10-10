package main

// type andmed struct{
//    eesnimi 		string     `form:text,valid:required`
//    perekonnanimi 	string     `form:text,valid:required`
//    email 			string     `form:text,valid:required|valid_email`
//    telefon 		int        `form:text,valid:required|numeric`
// }

import (
	"database/sql"
	"os"
	"github.com/lib/pq"
	"log"
	"net/http"
	"database/sql/driver"
	)

// start db connection
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

// check connection
// if err2 := database.Ping(); err2 != nil {
//  fmt.Println("Failed to keep connection alive")
// }

// get values from POST
func process_form_data(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		eesnimi := r.FormValue("eesnimi")
		perekonnanimi := r.FormValue("perekonnanimi")
		email := r.FormValue("email")
		telefon := r.FormValue("telefon")
	}
	result, err := database.Exec("INSERT INTO andmed(eesnimi,perekonnanimi,email,telefon) VALUES (eesnimi,perekonnanimi,email,telefon)")

}

// Exec func
func (db *DB) Exec(query string, args ...interface{}) (Result, error)

// insert to db
//result, err := database.Exec("INSERT INTO andmed(eesnimi,perekonnanimi,email,telefon) VALUES (eesnimi,perekonnanimi,email,telefon)")

// if err != nil {
//  log.Fatal(err)
// }
