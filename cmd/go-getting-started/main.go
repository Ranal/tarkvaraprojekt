package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	//"time"

	"github.com/ranal/tarkvaraprojekt/tarkvaraprojekt/Godeps/_workspace/src/github.com/gin-gonic/gin"
	_ "github.com/ranal/tarkvaraprojekt/tarkvaraprojekt/Godeps/_workspace/src/github.com/lib/pq"
	"github.com/ranal/tarkvaraprojekt/tarkvaraprojekt/Godeps/_workspace/src/github.com/russross/blackfriday"
)

var (
	repeat int
	db     *sql.DB = nil
)

func repeatFunc(c *gin.Context) {
	var buffer bytes.Buffer
	for i := 0; i < repeat; i++ {
		buffer.WriteString("Hello from Go!")
	}
	c.String(http.StatusOK, buffer.String())
}

func dbFunc(c *gin.Context) {
	/*
	if _, err := db.Exec("CREATE TABLE andmed(eesnimi varchar(50), perekonnanimi varchar(50), email varchar(50), telefon integer)"); err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error creating database table: %q", err))
		return
	}
	*/

	if _, err := db.Exec("INSERT INTO andmed VALUES ('Ranal', 'Saron', 'ranalsaron@gmail.com', 5120628)"); err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error incrementing tick: %q", err))
		return
	}

	//
		if _, err := db.Exec("INSERT INTO andmed VALUES ('Dino', 'Saurus', 'jurassic@gmail.com', 500999)"); err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error incrementing tick: %q", err))
		return
	}
	//

	rows, err := db.Query("SELECT eesnimi FROM andmed")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading ticks: %q", err))
		return
	}

	defer rows.Close()

	/*
	for rows.Next() {
		var eesnimi string
		if err := rows.Scan(&tick); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
	}
	*/
}

func main() {
	var err error
	var errd error
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tStr := os.Getenv("REPEAT")
	repeat, err = strconv.Atoi(tStr)
	if err != nil {
		log.Print("Error converting $REPEAT to an int: %q - Using default", err)
		repeat = 5
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

	router.GET("/mark", func(c *gin.Context) {
		c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("**hi!**"))))
	})

	router.GET("/repeat", repeatFunc)
	router.GET("/db", dbFunc)

	router.Run(":" + port)
}
