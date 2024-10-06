package routes

import (
	"backend/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
	{
		Uri:      "/users/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: true,
	},
	{
		Uri:      "/users",
		Method:   http.MethodGet,
		Function: controllers.FindAllUser,
		NeedAuth: true,
	},
	{
		Uri:      "/users/{id}",
		Method:   http.MethodGet,
		Function: controllers.FindUser,
		NeedAuth: true,
	},
	{
		Uri:      "/users/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: true,
	},
	{
		Uri:      "/users/{userId}/follow",
		Method:   http.MethodPost,
		Function: controllers.FollowUser,
		NeedAuth: true,
	},
	{
		Uri:      "/users/{userId}/unfollow",
		Method:   http.MethodPost,
		Function: controllers.UnfollowUser,
		NeedAuth: true,
	},
	{
		Uri:      "/users/{userId}/followers",
		Method:   http.MethodGet,
		Function: controllers.FindFollowers,
		NeedAuth: true,
	},
	{
		Uri:      "/users/{userId}/following",
		Method:   http.MethodGet,
		Function: controllers.FindFollowing,
		NeedAuth: true,
	},
	{
		Uri:      "/users/{userId}/update-password",
		Method:   http.MethodPost,
		Function: controllers.UpdatePassword,
		NeedAuth: true,
	},
}
