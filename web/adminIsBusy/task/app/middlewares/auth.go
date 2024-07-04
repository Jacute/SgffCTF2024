package middlewares

import (
	"KaifTravel/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token, err := c.Cookie("access_token")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		token, err := utils.CheckJwt(access_token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Incorrect access_token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			username := claims["username"].(string)
			c.Set("username", username)
			c.Set("isAuthenticated", true)
		}

		c.Next()
	}
}
