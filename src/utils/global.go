// Package utils provides utility functions for working with environment variables and templates.
package utils

import (
	"log"

	"os"

	"regexp"

	"github.com/joho/godotenv"
)

// Bt replaces placeholders in the template with the appropriate values.
// It processes ${env("SOMETHING")} placeholders by calling the Env function.
func Bt(template string) string {
	re := regexp.MustCompile(`\${env\("([^"]+)"\)}`)
	result := re.ReplaceAllStringFunc(template, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) >= 2 {
			return Env(parts[1])
		}

		return match
	})

	return result
}

// Env fetches the value of the specified environment variable.
func Env(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file!!")
	}

	return os.Getenv(key)
}
