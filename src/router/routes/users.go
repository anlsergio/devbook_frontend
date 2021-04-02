package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var usersRoutes = []Route{
	{
		URI: "/signup",
		Method: http.MethodGet,
		Function: controllers.RenderSignupPage,
		RequiresAuth: false,
	},
	{
		URI: "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		RequiresAuth: false,
	},
}