package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port         string
	Host         string
	Dsn          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func NewConfig() *Config {
	log.Println("Loading envs...")
	return &Config{
		Port:         GetEnv("PORT", ":8080"),
		Host:         GetEnv("HOST", "localhost"),
		Dsn:          GetEnv("DSN", "postgres://postgres:postgres@localhost/moraedb?sslmode=disable"),
		MaxOpenConns: GetEnvInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: GetEnvInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  GetEnv("DB_MAX_IDLE_TIME", "900s"),
	}
}

// Returns the value of an environment variable or a default if not set. Returns string.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

// Returns the value of an environment variable or a default if not set. Returns int.
func GetEnvInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if exists {
		conv, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}

		return conv
	}
	return fallback
}
