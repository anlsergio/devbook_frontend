package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"
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

// RenderHomePage - renders the home page listing user posts (feed)
func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	utils.RenderTemplate(w, "home.html", nil)
}