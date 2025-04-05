// Package env provides utility functions for reading and parsing environment variables
// with support for fallback values when variables are unset or invalid.
package env

import (
	"os"
	"strconv"
)

// GetBool retrieves a boolean env var or returns fallback if unset or invalid.
func GetBool(key string, fallback bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	parsed, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}

	return parsed
}

// GetString retrieves a string env var or returns fallback if unset.
func GetString(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	return val
}

// GetInt retrieves an integer env var or returns fallback if unset or invalid.
func GetInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	conv, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return conv
}
