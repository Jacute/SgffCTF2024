package config

import (
	"os"
)

type Config struct {
	JwtKey []byte
	Flag   string
}

var AppConfig = &Config{
	JwtKey: []byte(os.Getenv("JWT_KEY")),
	Flag:   os.Getenv("FLAG"),
}
