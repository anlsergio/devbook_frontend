package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var homePageRoute = Route{
	URI:    "/home",
	Method: http.MethodGet,
	Function: controllers.RenderHomePage,
	RequiresAuth: true,
}