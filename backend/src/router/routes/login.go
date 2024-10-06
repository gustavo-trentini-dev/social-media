package routes

import (
	"backend/src/controllers"
	"net/http"
)

var loginRoute = Route{
	Uri:      "/login",
	Method:   http.MethodPost,
	Function: controllers.Login,
	NeedAuth: false,
}
