package routes

import (
	"backend/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		Uri:      "/posts",
		Method:   http.MethodPost,
		Function: controllers.CreatePost,
		NeedAuth: true,
	},
	{
		Uri:      "/posts",
		Method:   http.MethodGet,
		Function: controllers.GetPosts,
		NeedAuth: true,
	},
	{
		Uri:      "/posts/{id}",
		Method:   http.MethodGet,
		Function: controllers.GetPost,
		NeedAuth: true,
	},
	{
		Uri:      "/posts/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdatePost,
		NeedAuth: true,
	},
	{
		Uri:      "/posts/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DeletePost,
		NeedAuth: true,
	},
	{
		Uri:      "/users/{userId}/posts",
		Method:   http.MethodGet,
		Function: controllers.FindUserPosts,
		NeedAuth: true,
	},
	{
		Uri:      "/posts/{id}/like",
		Method:   http.MethodPost,
		Function: controllers.LikePost,
		NeedAuth: true,
	},
	{
		Uri:      "/posts/{id}/dislike",
		Method:   http.MethodPost,
		Function: controllers.DislikePost,
		NeedAuth: true,
	},
}
