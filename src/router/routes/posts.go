package routes

import "webapp/src/controllers"

var postsRoutes = []Route {
	{
		URI:    "/posts",
		Method: "POST",
		Function: controllers.CreatePost,
		RequiresAuth: true,
	},
}