package middlewares

import (
	"errors"
	"net/http"
	"qrVachanaly/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func QrProgressMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token, err := c.Cookie("access_token")
		if err != nil {
			cookie, err := generateCookie()
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}

			http.SetCookie(c.Writer, cookie)
			c.Set("progress", 1)
			return
		}

		token, err := utils.CheckJwt(access_token)
		if err != nil {
			if errors.Is(utils.JwtExpiredError, err) {
				cookie, err := generateCookie()
				if err != nil {
					c.JSON(http.StatusOK, gin.H{"error": err.Error()})
					return
				}

				http.SetCookie(c.Writer, cookie)
				c.Set("progress", 1)
				return
			}
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

func generateCookie() (*http.Cookie, error) {
	new_access_token, err := utils.CreateJwt(1)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    new_access_token,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}
	return cookie, nil
}
