package handlers

import (
	"KaifTravel/database"
	"KaifTravel/models"
	"KaifTravel/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetLogin(c *gin.Context) {
	isAuthenticated, _ := utils.CheckAuthentication(c)

	if isAuthenticated {
		c.Redirect(http.StatusFound, "/")
		return
	}
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func PostLogin(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := database.GetUserByUsername(user.Username)
	if err != nil || !utils.CheckHash(user.Password, dbUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	jwt, err := utils.CreateJwt(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    jwt,
		Expires:  expirationTime,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(c.Writer, cookie)
	c.Redirect(http.StatusFound, "/")
}

func GetLogout(c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now(),
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}
	http.SetCookie(c.Writer, cookie)
	c.Redirect(http.StatusFound, "/")
}

func GetRegister(c *gin.Context) {
	isAuthenticated, _ := utils.CheckAuthentication(c)

	if isAuthenticated {
		c.Redirect(http.StatusFound, "/")
		return
	}
	c.HTML(200, "register.html", gin.H{})
}

func PostRegister(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userExists, err := database.UserExists(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	err = database.CreateUser(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jwt, err := utils.CreateJwt(user.Username)

	expirationTime := time.Now().Add(24 * time.Hour)
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    jwt,
		Expires:  expirationTime,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(c.Writer, cookie)
	c.Redirect(http.StatusFound, "/")
}
