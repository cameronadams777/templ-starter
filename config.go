package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(key string) string {
	// load .env file
	mode := os.Getenv("MODE")
	if mode != "release" {
		err := godotenv.Load(".env.development")
		if err != nil {
			fmt.Print("Error loading .env.development file")
		}
	}
	return os.Getenv(key)
}
