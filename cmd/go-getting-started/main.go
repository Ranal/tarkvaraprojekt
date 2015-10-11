package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

    //router.POST("/yhendus", func(c *gin.Context) {
   //}


    //-----//

    //func(c *gin.Context) {c.String(http.StatusOK, "see OK nupp funkab")}
	http.HandleFunc("/yhendus", yhendus(w http.ResponseWriter, req *http.Request))
    //-----//

	router.Run(":" + port)
}


//args: w http.ResponseWriter, req *http.Request
func yhendus(w http.ResponseWriter, req *http.Request) {
	//db, err := sql.Open("postgres", "user=vcjthhaofvkqke dbname=dedgfoiefjhcdu sslmode=disable 
	//	password=QXnZclsVqyZPU5C8Tn_ch81Qt2 host=ec2-54-217-238-100.eu-west-1.compute.amazonaws.com port=5432 ")
	db, err := sql.Open("postgres", "postgres://vcjthhaofvkqke:QXnZ	clsVqyZPU5C8Tn_ch81Qt2@ec2-54-217-238-100.eu-west-1.compute.amazonaws.com:5432/dedgfoiefjhcdu")
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