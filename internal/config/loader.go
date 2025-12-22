package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Get() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return &Config{
		Server: &Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: &Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			TimeZone: os.Getenv("DB_TIMEZONE"),
			Name:     os.Getenv("DB_NAME"),
		},
	}, nil
}
