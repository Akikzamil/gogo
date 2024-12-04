package utils

import (
	"os"
	"github.com/joho/godotenv"
)

func Getenv(key string, defaultValue string) string {
	err := godotenv.Load()

	if err != nil {
		return defaultValue;
	} else {
		env := os.Getenv(key)
		if env != "" {
			return env;
		}
	}
	return defaultValue;
}