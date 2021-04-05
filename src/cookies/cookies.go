package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Configure - makes use of the hash key and block key provided as env in order to generate a new cookie.
func Configure() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Set - stores the authentication information in the browser as a cookie
func Set(w http.ResponseWriter, ID, token string) error {
	data := map[string]string {
		"id": ID,
		"token": token,
	}

	encodedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name: "data",
		Value: encodedData,
		Path: "/",
		HttpOnly: true,
	})

	return nil
}

// Read - returns data from the cookie
func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	decodedData := make(map[string]string)
	if err = s.Decode("data", cookie.Value, &decodedData); err != nil {
		return nil, err
	}

	return decodedData, nil
}