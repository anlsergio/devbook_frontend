package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// RenderLoginPage - renders the login page in order to be loaded by the client's browser
func RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

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

	var posts []models.Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.RenderTemplate(w, "home.html", struct {
		Posts []models.Post
		UserID uint64
	}{
		Posts: posts,
		UserID: userID,
	})
}

// RenderEditPostPage - renders the post editing page
func RenderEditPostPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	endpoint_URL := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, endpoint_URL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	var post models.Post
	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	utils.RenderTemplate(w, "edit-post.html", post)
}


// RenderPageOfUsers - renders a new page containing a list of users who meet the query parameter
func RenderPageOfUsers(w http.ResponseWriter, r *http.Request) {
	nameOrUsername := strings.ToLower(r.URL.Query().Get("user"))

	endpoint_URL := fmt.Sprintf("%s/users?user=%s", config.APIURL, nameOrUsername)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, endpoint_URL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	utils.RenderTemplate(w, "users.html", users)
}