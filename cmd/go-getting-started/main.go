package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ranal/tarkvaraprojekt/tarkvaraprojekt/Godeps/_workspace/src/github.com/gin-gonic/gin"
	_ "github.com/ranal/tarkvaraprojekt/tarkvaraprojekt/Godeps/_workspace/src/github.com/lib/pq"
)

var (
	db     *sql.DB = nil
)

func dbFunc(c *gin.Context) {

	if _, err := db.Exec("INSERT INTO andmed VALUES ('Uus', 'Rida', 'uuedread@gmail.com', 5009208)"); err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error: %q", err))
		return
	}
	
	rows, err := db.Query("SELECT eesnimi FROM andmed")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading rows: %q", err))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("Rida andmebaasi lisatud"))

	defer rows.Close()
	
}

func main() {

	var errd error

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errd != nil {
		log.Fatalf("Error opening database: %q", errd)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.POST("/db", dbFunc)

	router.Run(":" + port)
}
