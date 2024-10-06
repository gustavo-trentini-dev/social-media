package routes

import (
	"frontend/src/controllers"
	"net/http"
)

var logoutRoutes = Route{
	URI:      "/logout",
	Method:   http.MethodGet,
	Func:     controllers.Logout,
	NeedAuth: true,
}
