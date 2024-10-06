package controllers

import (
	"backend/src/auth"
	"backend/src/database"
	"backend/src/models"
	"backend/src/repositories"
	"backend/src/response"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	loggedUserId, err := auth.GetUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post

	if err = json.Unmarshal(body, &post); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorId = loggedUserId

	if err = post.Prepare(); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepo := repositories.NewPostRepo(db)
	post.ID, err = postRepo.Create(post)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	loggedUserId, err := auth.GetUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepo := repositories.NewPostRepo(db)
	posts, err := postRepo.FindPosts(loggedUserId)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepo := repositories.NewPostRepo(db)
	post, err := postRepo.FindById(postID)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	loggedUserId, err := auth.GetUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepo := repositories.NewPostRepo(db)
	dbPost, err := postRepo.FindById(postID)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	if dbPost.AuthorId != loggedUserId {
		response.Err(w, http.StatusForbidden, errors.New("Not allowed to update other user post"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(body, &post); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = post.Prepare(); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = postRepo.Update(postID, post); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	loggedUserId, err := auth.GetUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepo := repositories.NewPostRepo(db)
	dbPost, err := postRepo.FindById(postID)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	if dbPost.AuthorId != loggedUserId {
		response.Err(w, http.StatusForbidden, errors.New("Not allowed to delete other user post"))
		return
	}

	if err = postRepo.Delete(postID); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func FindUserPosts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepo := repositories.NewPostRepo(db)
	posts, err := postRepo.FindUserPosts(userID)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepo := repositories.NewPostRepo(db)
	if err = postRepo.LikePost(postID); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DislikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepo := repositories.NewPostRepo(db)
	if err = postRepo.DislikePost(postID); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
