// Package config provides application configuration management by loading values from environment variables,
// with support for default fallbacks when variables are not set.
package config

import (
	"log"
	"os"
	"strconv"
)

// Config holds configuration values loaded from environment variables.
type Config struct {
	Port         string
	Host         string
	Dsn          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
	MongoDsn     string
}

// NewConfig initializes a new Config instance using environment variables or default values.
func NewConfig() *Config {
	log.Println("Loading envs...")
	return &Config{
		Port:         GetEnv("PORT", ":8080"),
		Host:         GetEnv("HOST", "localhost"),
		Dsn:          GetEnv("DSN", "postgres://postgres:postgres@localhost/moraedb?sslmode=disable"),
		MaxOpenConns: GetEnvInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: GetEnvInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  GetEnv("DB_MAX_IDLE_TIME", "900s"),
		MongoDsn:     GetEnv("MONGO_DSN", "mongodb://root:example@localhost:27017"),
	}
}

// GetEnv returns the value of an environment variable, or a fallback value if the variable is not set.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// GetEnvInt returns the integer value of an environment variable, or a fallback value if the variable is not set or conversion fails.
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
