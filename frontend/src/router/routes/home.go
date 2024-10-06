package routes

import (
	"frontend/src/controllers"
	"net/http"
)

var homeRoutes = Route{
	URI:      "/home",
	Method:   http.MethodGet,
	Func:     controllers.LoadHome,
	NeedAuth: true,
}
