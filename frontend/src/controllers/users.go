package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend/src/config"
	"frontend/src/cookies"
	"frontend/src/requests"
	"frontend/src/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		response.JSON(w, http.StatusBadRequest,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}

	url := fmt.Sprintf("%s/users", config.API_URL)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(user))
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

func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorApi{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.API_URL, userID)
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

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorApi{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.API_URL, userID)
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

func EditProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
	})

	if err != nil {
		response.JSON(w, http.StatusBadRequest,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.API_URL, userID)
	resp, err := requests.DoRequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(user))
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

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	passwords, err := json.Marshal(map[string]string{
		"new":     r.FormValue("new"),
		"current": r.FormValue("current"),
	})

	if err != nil {
		response.JSON(w, http.StatusBadRequest,
			response.ErrorApi{
				Error: err.Error(),
			},
		)
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/update-password", config.API_URL, userID)
	resp, err := requests.DoRequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(passwords))
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.API_URL, userID)
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
