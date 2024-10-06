package controllers

import (
	"backend/src/auth"
	"backend/src/database"
	"backend/src/models"
	"backend/src/repositories"
	"backend/src/response"
	"backend/src/secure"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUserRepo(db)
	userFromDb, err := userRepo.FindByEmail(user.Email)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
	}

	if err = secure.VerifyPass(userFromDb.Password, user.Password); err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(userFromDb.ID)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(userFromDb.ID, 10)

	response.JSON(w, http.StatusOK, models.Auth{
		ID:    userID,
		Token: token,
	})
}
