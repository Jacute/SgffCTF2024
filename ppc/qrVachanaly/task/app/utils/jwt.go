package utils

import (
	"errors"
	"qrVachanaly/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
	JwtError        = errors.New("JWT error")
	JwtExpiredError = errors.New("JWT expired")
)

func CheckJwt(access_token string) (*jwt.Token, error) {
	if access_token == "" {
		return nil, JwtError
	}

	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		return config.AppConfig.JwtKey, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && int64(claims["exp"].(float64)) < time.Now().Unix() {
		return nil, JwtExpiredError
	}
	if err != nil || !token.Valid {
		return nil, err
	}
	return token, nil
}

func CreateJwt(progress int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = uuid.New().String()
	claims["progress"] = progress
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	tokenString, err := token.SignedString(config.AppConfig.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
