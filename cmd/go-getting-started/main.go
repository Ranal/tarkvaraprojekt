package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {

	//
	http.HandleFunc("/yhendus",andmedFunc)
	http.ListenAndServe(":8080", nil)
	//

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

	router.Run(":" + port)
}

func andmedFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}