package routes

import (
	"frontend/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Func     func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}

func ConfigRoutes(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, usersRoutes...)
	routes = append(routes, postsRoutes...)
	routes = append(routes, homeRoutes)
	routes = append(routes, logoutRoutes)

	for _, route := range routes {

		if route.NeedAuth {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Auth(route.Func)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(route.Func),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
