package routes

import (
	"frontend/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:      "/create-user",
		Method:   http.MethodGet,
		Func:     controllers.LoadRegisterUser,
		NeedAuth: false,
	},
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Func:     controllers.CreateUser,
		NeedAuth: false,
	},
	{
		URI:      "/search-users",
		Method:   http.MethodGet,
		Func:     controllers.LoadUsersPage,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodGet,
		Func:     controllers.LoadUserProfile,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/follow",
		Method:   http.MethodPost,
		Func:     controllers.FollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/unfollow",
		Method:   http.MethodPost,
		Func:     controllers.UnfollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/profile",
		Method:   http.MethodGet,
		Func:     controllers.LoadLoggedUserProfile,
		NeedAuth: true,
	},
	{
		URI:      "/edit-user",
		Method:   http.MethodGet,
		Func:     controllers.LoadEditProfile,
		NeedAuth: true,
	},
	{
		URI:      "/edit-user",
		Method:   http.MethodPut,
		Func:     controllers.EditProfile,
		NeedAuth: true,
	},
	{
		URI:      "/update-password",
		Method:   http.MethodGet,
		Func:     controllers.LoadUpdatePassword,
		NeedAuth: true,
	},
	{
		URI:      "/update-password",
		Method:   http.MethodPut,
		Func:     controllers.UpdatePassword,
		NeedAuth: true,
	},
	{
		URI:      "/delete-user",
		Method:   http.MethodDelete,
		Func:     controllers.DeleteUser,
		NeedAuth: true,
	},
}
