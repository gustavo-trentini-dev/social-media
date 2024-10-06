package routes

import (
	"frontend/src/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:      "/",
		Method:   http.MethodGet,
		Func:     controllers.LoadLogin,
		NeedAuth: false,
	},
	{
		URI:      "/login",
		Method:   http.MethodGet,
		Func:     controllers.LoadLogin,
		NeedAuth: false,
	},
	{
		URI:      "/login",
		Method:   http.MethodPost,
		Func:     controllers.Login,
		NeedAuth: false,
	},
}
