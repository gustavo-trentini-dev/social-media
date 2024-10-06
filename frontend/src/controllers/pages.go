package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/src/config"
	"frontend/src/cookies"
	"frontend/src/models"
	"frontend/src/requests"
	"frontend/src/response"
	"frontend/src/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func LoadLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
	}

	utils.ExecTemplate(w, "login.html", nil)
}

func LoadRegisterUser(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "register.html", nil)
}

func LoadHome(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.API_URL)
	resp, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
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

	var posts []models.Post
	if err = json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserId uint64
	}{
		Posts:  posts,
		UserId: userID,
	})
}

func LoadEditPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorApi{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.API_URL, postID)
	resp, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
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

	var post models.Post
	if err = json.NewDecoder(resp.Body).Decode(&post); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorApi{Error: err.Error()})
		return
	}

	utils.ExecTemplate(w, "update-post.html", post)
}

func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	url := fmt.Sprintf("%s/users?user=%s", config.API_URL, nameOrNick)
	resp, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
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

	var users []models.User
	if err = json.NewDecoder(resp.Body).Decode(&users); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorApi{Error: err.Error()})
		return
	}

	utils.ExecTemplate(w, "users.html", users)
}

func LoadUserProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorApi{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userLoggedID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userID == userLoggedID {
		http.Redirect(w, r, "/profile", 302)
		return
	}

	user, err := models.GetFullUser(userID, r)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorApi{Error: err.Error()})
		return
	}

	utils.ExecTemplate(w, "user.html", struct {
		User         models.User
		UserLoggedId uint64
	}{
		User:         user,
		UserLoggedId: userLoggedID,
	})
}

func LoadLoggedUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.GetFullUser(userID, r)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorApi{Error: err.Error()})
		return
	}

	utils.ExecTemplate(w, "profile.html", user)
}

func LoadEditProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan models.User)
	go models.GetUser(channel, userID, r)
	user := <-channel

	if user.ID == 0 {
		response.JSON(w, http.StatusInternalServerError, response.ErrorApi{Error: "Error searching user"})
		return
	}

	utils.ExecTemplate(w, "edit-user.html", user)
}

func LoadUpdatePassword(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "update-password.html", nil)
}
