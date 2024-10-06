package main

import (
	"fmt"
	"frontend/src/config"
	"frontend/src/cookies"
	"frontend/src/router"
	"frontend/src/utils"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	cookies.Config()
	utils.LoadTemplates()
	r := router.Generate()

	fmt.Printf("Running frontend %d", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))
}
