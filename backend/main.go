package main

import (
	"backend/src/config"
	"backend/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	r := router.Generate()

	fmt.Printf("Listening on port: %d", config.API_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), r))
}
