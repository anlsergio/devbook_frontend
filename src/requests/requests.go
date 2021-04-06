package requests

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// RequestWithAuthentication is used to create a new request in order to call the API with the JWT token injected in its header
func RequestWithAuthentication(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	// Create a new request in order to call the API
	newRequest, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	// Fetch the cookie data from the initial request which carries the valid token (created initially by the frontend)
	cookie, _ := cookies.Read(r)
	newRequest.Header.Add("Authorization", "Bearer "+ cookie["token"])

	client := &http.Client{}
	response, err := client.Do(newRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}