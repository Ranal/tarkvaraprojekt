package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"bytes"
	"strconv"
	"github.com/russross/blackfriday"
)

func andmedFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

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

	router.GET("/yhendus", func(c *gin.Context) {
        c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("**hi!**"))))
    })

    router.GET("/repeat", andmedFunc)

	router.Run(":" + port)

//
	http.HandleFunc("/yhendus",func(w http.ResponseWriter, req *http.Request){w.Write([]byte("Hello World"))})
	http.ListenAndServe(":"+port, nil)

//
}

