package main

import (
	"log"
	"net/http"
	"qrVachanaly/handlers"
	"qrVachanaly/middlewares"
	"qrVachanaly/qrGenerator"

	"github.com/gin-gonic/gin"
)

func main() {
	go qrGenerator.Start()

	router := gin.Default()

	router.LoadHTMLGlob("./templates/*.html")
	router.Static("/static", "./static")
	router.Static("/qrs", "./qrs")

	router.GET("/", middlewares.QrProgressMiddleware(), handlers.Index)
	router.POST("/", middlewares.QrProgressMiddleware(), handlers.PostIndex)

	log.Fatal(http.ListenAndServe(":8080", router))
}
