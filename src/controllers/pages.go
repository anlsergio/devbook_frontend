package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// RenderLoginPage - renders the login page in order to be loaded by the client's browser
func RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login.html", nil)
}

// RenderSignupPage - renders the users' sign up page
func RenderSignupPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "signup.html", nil)
}