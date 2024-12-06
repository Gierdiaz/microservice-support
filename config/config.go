package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Database struct {
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
}

type Server struct {
	APP_PORT string
}

type JWT struct {
	SECRET_KEY   string
	TOKEN_EXPIRE int
}

type Config struct {
	Database Database
	JWT      JWT
	Server   Server
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	tokenExpireStr  := os.Getenv("TOKEN_EXPIRE")
	tokenExpire := 3600
	if tokenExpireStr != "" {
		parsed, err := strconv.Atoi(tokenExpireStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing TOKEN_EXPIRE: %v", err)
		}
		tokenExpire = parsed
	}

	config := &Config{
		Server: Server{
			APP_PORT: os.Getenv("APP_PORT"),
		},
		Database: Database{
			DB_HOST:     os.Getenv("DB_HOST"),
			DB_PORT:     os.Getenv("DB_PORT"),
			DB_USERNAME: os.Getenv("DB_USERNAME"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
			DB_DATABASE: os.Getenv("DB_DATABASE"),
		},
		JWT: JWT{
			SECRET_KEY:   os.Getenv("SECRET_KEY"),
			TOKEN_EXPIRE: tokenExpire,
		},
	}

	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("error validating config: %v", err)
	}

	return config, nil
}

func validateConfig(config *Config) error {
	if config.Server.APP_PORT == "" {
		return fmt.Errorf("APP_PORT is not set")
	}

	if config.Database.DB_HOST == "" {
		return fmt.Errorf("DB_HOST is not set")
	}

	if config.Database.DB_PORT == "" {
		return fmt.Errorf("DB_PORT is not set")
	}

	if config.Database.DB_USERNAME == "" {
		return fmt.Errorf("DB_USERNAME is not set")
	}

	if config.Database.DB_PASSWORD == "" {
		return fmt.Errorf("DB_PASSWORD is not set")
	}

	if config.Database.DB_DATABASE == "" {
		return fmt.Errorf("DB_DATABASE is not set")
	}

	if config.JWT.SECRET_KEY == "" {
		return fmt.Errorf("SECRET_KEY is not set")
	}

	return nil
}