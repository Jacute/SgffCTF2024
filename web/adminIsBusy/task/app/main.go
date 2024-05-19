package main

import (
	"KaifTravel/database"
	"KaifTravel/handlers"
	"KaifTravel/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var JwtKey string = "secretkey"

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*.html")
	router.Static("/static", "./static")

	router.GET("/", middlewares.AuthMiddleware(), handlers.GetIndex)
	router.GET("/login", handlers.GetLogin)
	router.POST("/login", handlers.PostLogin)
	router.GET("/register", handlers.GetRegister)
	router.POST("/register", handlers.PostRegister)
	router.GET("/logout", handlers.GetLogout)
	db := database.InitSqliteDB()
	defer db.Close()

	database.CreateDB()

	log.Fatal(http.ListenAndServe(":8081", router))
}
