package main

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	port         string
	host         string
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func newConfig() *Config {
	log.Println("Loading envs...")
	return &Config{
		port:         getEnv("PORT", ":8080"),
		host:         getEnv("HOST", "localhost"),
		dsn:          getEnv("DSN", "postgres://victorsvart:123qwe@localhost/moraedb?sslmode=disable"),
		maxOpenConns: getEnvInt("DB_MAX_OPEN_CONNS", 30),
		maxIdleConns: getEnvInt("DB_MAX_IDLE_CONNS", 30),
		maxIdleTime:  getEnv("DB_MAX_IDLE_TIME", "900s"),
	}
}

// Returns the value of an environment variable or a default if not set. Returns string.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

// Returns the value of an environment variable or a default if not set. Returns int.
func getEnvInt(key string, fallback int) int {
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
