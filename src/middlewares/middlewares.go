package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger - Prints out information regarding the current request
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// CheckAuth - checks if the user making the request is currently authenticated
func CheckAuth(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookies.Read(r); err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		nextFunc(w, r)
	}
}