package controllers

import (
	"backend/src/auth"
	"backend/src/database"
	"backend/src/models"
	"backend/src/repositories"
	"backend/src/response"
	"backend/src/secure"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.Prepare("register"); err != nil {
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

	user.ID, err = userRepo.Create(user)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	loggedUserId, err := auth.GetUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	if userID != loggedUserId {
		response.Err(w, http.StatusForbidden, errors.New("You can only update your own user"))
		return
	}

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

	if err = user.Prepare("update"); err != nil {
		response.Err(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUserRepo(db)

	if err = userRepo.Update(userID, user); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func FindAllUser(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUserRepo(db)

	users, err := userRepo.FindAll(nameOrNick)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
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

	userRepo := repositories.NewUserRepo(db)

	user, err := userRepo.FindOne(userID)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	loggedUserId, err := auth.GetUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	if userID != loggedUserId {
		response.Err(w, http.StatusForbidden, errors.New("You can only delete your own user"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUserRepo(db)

	if err = userRepo.Delete(userID); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.GetUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		response.Err(w, http.StatusForbidden, errors.New("You cannot follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUserRepo(db)

	if err = userRepo.Follow(userId, followerId); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.GetUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		response.Err(w, http.StatusForbidden, errors.New("You cannot unfollow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUserRepo(db)

	if err = userRepo.Unfollow(userId, followerId); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func FindFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
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

	userRepo := repositories.NewUserRepo(db)
	followers, err := userRepo.FindFollowers(userId)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, followers)

}

func FindFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
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

	userRepo := repositories.NewUserRepo(db)
	followers, err := userRepo.FindFollowing(userId)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	loggerUserID, err := auth.GetUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if loggerUserID != userId {
		response.Err(w, http.StatusForbidden, errors.New("Not possible to update password from other user"))
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password models.Password

	if err = json.Unmarshal(body, &password); err != nil {
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
	dbPass, err := userRepo.FindById(userId)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = secure.VerifyPass(dbPass, password.Current); err != nil {
		response.Err(w, http.StatusUnauthorized, errors.New("The current password doesn't match"))
		return
	}

	newHashPass, err := secure.Hash(password.New)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = userRepo.UpdatePassword(userId, string(newHashPass)); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
