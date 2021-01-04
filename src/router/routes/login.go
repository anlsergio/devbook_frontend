package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var loginRoutes = []Route{
	{
		URI:    "/",
		Method: http.MethodGet,
		Function: controllers.RenderLoginPage,
		RequiresAuth: false,
	},
	{
		URI:    "/login",
		Method: http.MethodGet,
		Function: controllers.RenderLoginPage,
		RequiresAuth: false,
	},
}