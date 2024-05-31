package config

import (
	"os"

	"github.com/google/uuid"
)

type Config struct {
	JwtKey []byte
	Flag   string
}

var AppConfig = &Config{
	JwtKey: []byte(uuid.New().String()),
	Flag:   os.Getenv("FLAG"),
}
