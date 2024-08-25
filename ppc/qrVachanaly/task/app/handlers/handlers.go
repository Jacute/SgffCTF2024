package handlers

import (
	"net/http"
	"qrVachanaly/config"
	"qrVachanaly/qrGenerator"
	"qrVachanaly/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	progress, exists := c.Get("progress")
	if progress == 101 {
		c.Data(200, "text/html; charset=utf-8", []byte(config.AppConfig.Flag))
		return
	}

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Progress not found"})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"progress": progress})
}

func PostIndex(c *gin.Context) {
	progress, exists := c.Get("progress")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Field progress not exists in jwt"})
		return
	}

	progressInt, ok := progress.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot convert progress to int"})
		return
	}

	answer := c.PostForm("answer")
	if answer == qrGenerator.QrData[progressInt] {
		new_access_token, err := utils.CreateJwt(progressInt + 1)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't create jwt"})
			return
		}

		expirationTime := time.Now().Add(24 * time.Hour)
		cookie := &http.Cookie{
			Name:     "access_token",
			Value:    new_access_token,
			Expires:  expirationTime,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteDefaultMode,
		}

		http.SetCookie(c.Writer, cookie)

		c.JSON(http.StatusOK, gin.H{"status": "SUCCESS"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid answer"})
	}
}
