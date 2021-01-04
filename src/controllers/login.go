package controllers

import "net/http"

// RenderLoginPage - renders the login page in order to be loaded by the client's browser
func RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login page"))
}