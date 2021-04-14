package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postsRoutes = []Route {
	{
		URI:    "/posts",
		Method: http.MethodPost,
		Function: controllers.CreatePost,
		RequiresAuth: true,
	},
	{
		URI:    "/posts/{postID}/like",
		Method: http.MethodPost,
		Function: controllers.LikePost,
		RequiresAuth: true,
	},
	{
		URI:    "/posts/{postID}/dislike",
		Method: http.MethodPost,
		Function: controllers.DislikePost,
		RequiresAuth: true,
	},
	{
		URI:    "/posts/{postID}/edit",
		Method: http.MethodGet,
		Function: controllers.RenderEditPostPage,
		RequiresAuth: true,
	},
	{
		URI:    "/posts/{postID}",
		Method: http.MethodPut,
		Function: controllers.UpdatePost,
		RequiresAuth: true,
	},
}