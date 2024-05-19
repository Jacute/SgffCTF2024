package utils

import (
	"KaifTravel/config"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func RandomPassword(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateJwt(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(config.AppConfig.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckJwt(access_token string) (*jwt.Token, error) {
	if access_token == "" {
		return nil, &CustomError{message: "Token is empty string"}
	}

	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		return config.AppConfig.JwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token, nil
}

func CheckAuthentication(c *gin.Context) (bool, error) {
	access_token, err := c.Cookie("access_token")
	if err != nil {
		return false, err
	}
	if access_token == "" {
		return false, &CustomError{message: "Token is empty"}
	}

	_, err = CheckJwt(access_token)
	if err != nil {
		return false, nil
	}
	return true, nil
}
