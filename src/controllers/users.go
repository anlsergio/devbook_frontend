package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

// CreateUser - calls the API's endpoint for creating a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name": r.FormValue("name"),
		"email": r.FormValue("email"),
		"username": r.FormValue("username"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	endpoint_URL := fmt.Sprintf("%s/users", config.APIURL)
	response, err := http.Post(endpoint_URL, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// FollowUser - calls the API's endpoint in order to follow a user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	endpoint_URL := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userID)
	response, err := requests.RequestWithAuthentication(r, http.MethodPost, endpoint_URL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UnfollowUser - calls the API's endpoint in order to follow a user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	endpoint_URL := fmt.Sprintf("%s/users/%d/unfollow", config.APIURL, userID)
	response, err := requests.RequestWithAuthentication(r, http.MethodDelete, endpoint_URL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UpdateUser - calls the API's endpoint in order to update a user's personal information
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name": r.FormValue("name"),
		"email": r.FormValue("email"),
		"username": r.FormValue("username"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	endpoint_URL := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, err := requests.RequestWithAuthentication(r, http.MethodPut, endpoint_URL, bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}