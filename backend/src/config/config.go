package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DB_CONNECTION = ""
	API_PORT      = 0
	SECRET_KEY    []byte
)

func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	API_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		API_PORT = 5000
	}

	PORT, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		PORT = 5432
	}

	DB_CONNECTION = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		PORT,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

}
