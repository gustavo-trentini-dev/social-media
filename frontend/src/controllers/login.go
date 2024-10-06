package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend/src/config"
	"frontend/src/cookies"
	"frontend/src/models"
	"frontend/src/response"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
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

	url := fmt.Sprintf("%s/login", config.API_URL)
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

	var auth models.Auth

	if err = json.NewDecoder(resp.Body).Decode(&auth); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity,
			response.ErrorApi{
				Error: err.Error(),
			})
		return
	}

	if err = cookies.Save(w, auth.ID, auth.Token); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity,
			response.ErrorApi{
				Error: err.Error(),
			})
		return
	}

	response.JSON(w, resp.StatusCode, nil)
}
