package controllers

import (
	"frontend/src/cookies"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", 302)
}
