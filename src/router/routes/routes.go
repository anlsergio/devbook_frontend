package routes

import (
	"net/http"
	"webapp/src/middlewares"

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
	routes = append(routes, usersRoutes...)
	routes = append(routes, homePageRoute)
	routes = append(routes, postsRoutes...)
	routes = append(routes, logoutRoute)

	for _, route := range routes {
		if route.RequiresAuth {
			router.HandleFunc(route.URI, 
				middlewares.Logger(middlewares.CheckAuth(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}

		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}