package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	API_URL   = ""
	PORT      = 0
	HASH_KEY  []byte
	BLOCK_KEY []byte
)

func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		PORT = 3000
	}

	API_URL = os.Getenv("API_URL")
	HASH_KEY = []byte(os.Getenv("HASH_KEY"))
	BLOCK_KEY = []byte(os.Getenv("BLOCK_KEY"))
}
