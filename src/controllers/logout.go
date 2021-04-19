package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Logout - Logs the user out by invalidating the token data from the browser
func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Invalidate(w)
	http.Redirect(w, r, "/login", 302)
}