package config

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

// Config структура описания конфигурации
type Config struct {
	HttpPort string
	StartDir string
}

// NewConfig создаем конфигурацию приложения
func NewConfig(path string) *Config {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("no .env file, reading config from OS ENV variables")
	}
	return &Config{
		HttpPort: os.Getenv("HTTP_PORT"),
		StartDir: path,
	}
}
