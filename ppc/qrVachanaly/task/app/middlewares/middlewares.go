package middlewares

import (
	"net/http"
	"qrVachanaly/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func QrProgressMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token, err := c.Cookie("access_token")
		if err != nil {
			new_access_token, err := utils.CreateJwt(1)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}

			expirationTime := time.Now().Add(15 * time.Minute)
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
			c.Set("progress", 1)
			return
		}

		token, err := utils.CheckJwt(access_token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid jwt"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			progress := claims["progress"].(float64)
			c.Set("progress", int(progress))
		}

		c.Next()
	}
}
