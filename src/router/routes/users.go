package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:          "/signup",
		Method:       http.MethodGet,
		Function:     controllers.RenderSignupPage,
		RequiresAuth: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/search-users",
		Method:       http.MethodGet,
		Function:     controllers.RenderPageOfUsers,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodGet,
		Function:     controllers.RenderUsersProfile,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userID}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userID}/unfollow",
		Method:       http.MethodDelete,
		Function:     controllers.UnfollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/profile",
		Method:       http.MethodGet,
		Function:     controllers.RenderSignedInUserProfile,
		RequiresAuth: true,
	},
	{
		URI:          "/edit-user",
		Method:       http.MethodGet,
		Function:     controllers.RenderEditUserPage,
		RequiresAuth: true,
	},
	{
		URI:          "/update-user",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiresAuth: true,
	},
}
