package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responses"
)

// Login - calls the API in order to retrieve authentication information regarding a given user
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string {
		"email": r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}

	endpoint_URL := fmt.Sprintf("%s/login", config.APIURL)
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

	var authData models.AuthData
	if err = json.NewDecoder(response.Body).Decode(&authData); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	// cookie creation
	if err = cookies.Set(w, authData.ID, authData.Token); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	responses.JSON(w, http.StatusOK, nil)

}