package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBUrl    string
	ChainUrl string
)

func Config() error {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("loading env: ", err)
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return errors.New("DB_URL cannot be empty")
	}
	DBUrl = dbUrl

	chainUrl := os.Getenv("BEACON_CHAIN_URL")
	if chainUrl == "" {
		return errors.New("BEACON_CHAIN_URL cannot be empty")
	}
	ChainUrl = chainUrl

	return nil
}
