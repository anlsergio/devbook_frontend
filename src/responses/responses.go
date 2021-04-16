package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIError - represents the error reponse from the API
type APIError struct {
	Error string `json:"error"`
}

// JSON - returns a JSON as a response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// HandleErrorStatusCode - handles requests with a status code of 400 and up
func HandleErrorStatusCode(w http.ResponseWriter, r *http.Response) {
	var error APIError
	json.NewDecoder(r.Body).Decode(&error)
	JSON(w, r.StatusCode, error)
}