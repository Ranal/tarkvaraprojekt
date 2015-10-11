package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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


func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }

    //
	fmt.Printf("%+v\n", r.Form)
	for key, values := range r.Form {   // range over map
	  for _, value := range values {    // range over []string
	     fmt.Println(key, value)
	  }
	}

	//
    //fmt.Fprintf(w, "Hello world!") // write data to response
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

	router.GET("/db_", dbFunc)
	router.GET("/db", func(c *gin.Context){
		sayhelloName(c.Writer, c.Request)
		})
	//http.HandleFunc("/db", sayhelloName)

	router.Run(":" + port)
}
