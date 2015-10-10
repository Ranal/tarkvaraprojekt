package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	//db, err := sql.Open("postgres", "user=vcjthhaofvkqke dbname=dedgfoiefjhcdu sslmode=disable 
	//	password=QXnZclsVqyZPU5C8Tn_ch81Qt2 host=ec2-54-217-238-100.eu-west-1.compute.amazonaws.com port=5432 ")
	
db, err := sql.Open("postgres", "postgres://vcjthhaofvkqke:QXnZclsVqyZPU5C8Tn_ch81Qt2@ec2-54-217-238-100.eu-west-1.compute.amazonaws.com:5432/dedgfoiefjhcdu")
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
}