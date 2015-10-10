package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)


func yhendus(w http.ResponseWriter, req *http.Request) {
	//db, err := sql.Open("postgres", "user=vcjthhaofvkqke dbname=dedgfoiefjhcdu sslmode=disable 
	//	password=QXnZclsVqyZPU5C8Tn_ch81Qt2 host=ec2-54-217-238-100.eu-west-1.compute.amazonaws.com port=5432 ")
	db, err := sql.Open("postgres", "postgres://vcjthhaofvkqke:QXnZclsVqyZPU5C8Tn_ch81Qt2@ec2-54-217-238-100.eu-west-1.compute.amazonaws.com:5432/dedgfoiefjhcdu")
	if err != nil {
		log.Fatal(err)
	}
stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
if err != nil {
	log.Fatal(err)
}
res, err := stmt.Exec("Dolly")
if err != nil {
	log.Fatal(err)
}
lastId, err := res.LastInsertId()
if err != nil {
	log.Fatal(err)
}
rowCnt, err := res.RowsAffected()
if err != nil {
	log.Fatal(err)
}
log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func main() {
	http.HandleFunc("/yhendus", yhendus)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}