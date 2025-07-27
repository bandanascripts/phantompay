package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load environment variable from env file : %v", err)
	}
}
