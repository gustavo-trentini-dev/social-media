package router

import (
	"frontend/src/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.ConfigRoutes(r)
}
