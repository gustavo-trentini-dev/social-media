package routes

import (
	"frontend/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:      "/create-post",
		Method:   http.MethodPost,
		Func:     controllers.CreatePost,
		NeedAuth: true,
	},
	{
		URI:      "/update-post/{postId}",
		Method:   http.MethodPut,
		Func:     controllers.UpdatePost,
		NeedAuth: true,
	},
	{
		URI:      "/delete-post/{postId}",
		Method:   http.MethodDelete,
		Func:     controllers.DeletePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{postId}/like",
		Method:   http.MethodPost,
		Func:     controllers.LikePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{postId}/dislike",
		Method:   http.MethodPost,
		Func:     controllers.DislikePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{postId}/edit",
		Method:   http.MethodGet,
		Func:     controllers.LoadEditPost,
		NeedAuth: true,
	},
}
