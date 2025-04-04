package env

import (
	"os"
	"strconv"
)

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

func GetString(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	return val
}

func GetInt(key, fallback int) int {
	val := os.Getenv(strconv.Itoa(key))
	if val == "" {
		return fallback
	}

	conv, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return conv
}
