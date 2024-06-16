package main

import (
	"crc32_checker/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Mkdir("uploads", 0755)

	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("./views/*.html")

	router.GET("/", handlers.GetIndex)
	router.POST("/check", handlers.PostCheck)

	log.Fatal(http.ListenAndServe(":9999", router))
}
