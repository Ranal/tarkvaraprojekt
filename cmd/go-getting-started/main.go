package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	db     *sql.DB = nil
)

func dbFunc(c *gin.Context) {

	_, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS andmed(eesnimi varchar(50), perekonnanimi varchar(50), email varchar(50), telefon integer);
		CREATE TABLE IF NOT EXISTS sooduskoodid(kood varchar(50), soodustus float)",)
	if err != nil {
	 log.Fatal(err)
	}

	eesnimi := c.Query("eesnimi")
	perekonnanimi := c.Query("perekonnanimi")
	email := c.Query("email")
	telefon := c.Query("telefon")
	sooduskood := c.Query("sooduskood")

	c.String(http.StatusOK, "Pilet ostetud! Nimi: %s %s | E-mail: %s | Telefon: %s ", eesnimi, perekonnanimi, email, telefon)

	//todo
	if _, err := db.Exec("SELECT COUNT(1) FROM sooduskoodid WHERE kood = $1", sooduskood); err != nil {
	c.String(http.StatusInternalServerError,
		fmt.Sprintf("Sooduskoodi ei leitud! %q", err))
	return
	}

	if _, err := db.Exec("INSERT INTO andmed VALUES ($1, $2, $3, $4)", eesnimi, perekonnanimi, email, telefon); err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Viga andmete sisestamisel: %q", err))
		return
	}
	
	rows, err := db.Query("SELECT eesnimi FROM andmed")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Viga ridade lugemisel: %q", err))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("| Rida edukalt lisatud andmebaasi"))

	defer rows.Close()
}

func main() {

	var errd error

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT seadmata")
	}

	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errd != nil {
		log.Fatalf("Viga andmebaasiga connectimisel: %q", errd)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/db", dbFunc)

	router.Run(":" + port)
}
