package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL is the URL where the backend API is available
	APIURL = ""

	// Port is the Web Server port where the Web App application should be listening for connections
	Port = 0

	// HashKey is used to authenticate the cookie itself
	HashKey []byte

	// BlockKey is used to encrypt the cookie data
	BlockKey []byte
)

// Load initializes the environment variables for the application
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}