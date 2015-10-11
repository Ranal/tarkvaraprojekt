package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

var (
    db *sql.DB
)

func main() {

	db = newDb(DbConnection)

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

    router.POST("/yhendus", func(c *gin.Context) {
        c.String(http.StatusOK, "see OK nupp funkab")
    })

    //
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
	//

	router.Run(":" + port)
}

func newDb(connection string) *sql.DB {
	db, err := sql.Open("postgres", "postgres://vcjthhaofvkqke:QXnZclsVqyZPU5C8Tn_ch81Qt2@ec2-54-217-238-100.eu-west-1.compute.amazonaws.com:5432/dedgfoiefjhcdu")
    if err != nil {
        panic(err)
    }

    return db
}