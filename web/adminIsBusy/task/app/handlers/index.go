package handlers

import (
	"KaifTravel/config"
	"KaifTravel/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	isAuthenticated, _ := utils.CheckAuthentication(c)

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get username"})
		return
	}

	if username == "admin" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"username":        username,
			"flag":            config.AppConfig.Flag,
			"isAuthenticated": isAuthenticated,
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"username":        username,
		"isAuthenticated": isAuthenticated,
	})
}
