package main

import (
	"easyPeasy/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*.html")
	router.Static("/static", "./static")

	router.GET("/", handlers.GetIndex)
	router.GET("/part1", handlers.GetSecret1)
	router.GET("/part2", handlers.GetSecret2)

	log.Fatal(http.ListenAndServe(":8083", router))
}
