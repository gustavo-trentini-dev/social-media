package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend/src/config"
	"frontend/src/requests"
	"frontend/src/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		response.JSON(w, http.StatusBadRequest,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}

	url := fmt.Sprintf("%s/posts", config.API_URL)
	resp, err := requests.DoRequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(post))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		response.HandleErrorStatusCode(w, resp)
		return
	}

	response.JSON(w, resp.StatusCode, nil)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorApi{Error: err.Error()})
		return
	}

	r.ParseForm()
	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		response.JSON(w, http.StatusBadRequest,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.API_URL, postID)
	resp, err := requests.DoRequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(post))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		response.HandleErrorStatusCode(w, resp)
		return
	}

	response.JSON(w, resp.StatusCode, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorApi{Error: err.Error()})
		return
	}

	if err != nil {
		response.JSON(w, http.StatusBadRequest,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.API_URL, postID)
	resp, err := requests.DoRequestWithAuth(r, http.MethodDelete, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		response.HandleErrorStatusCode(w, resp)
		return
	}

	response.JSON(w, resp.StatusCode, nil)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorApi{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/like", config.API_URL, postID)
	resp, err := requests.DoRequestWithAuth(r, http.MethodPost, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		response.HandleErrorStatusCode(w, resp)
		return
	}

	response.JSON(w, resp.StatusCode, nil)
}

func DislikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorApi{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/dislike", config.API_URL, postID)
	resp, err := requests.DoRequestWithAuth(r, http.MethodPost, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		response.HandleErrorStatusCode(w, resp)
		return
	}

	response.JSON(w, resp.StatusCode, nil)
}
