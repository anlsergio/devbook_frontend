package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all API routes
type Route struct {
	URI 					string
	Method 				string
	Function 			func(http.ResponseWriter, *http.Request)
	RequiresAuth	bool
}

// SetRouter configures all routes inside the router
func SetRouter(router *mux.Router) *mux.Router {
	routes := loginRoutes

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router
}