package middleware

import (
	"net/http"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from reuquest
		next.ServeHTTP(w, r)
	})
}
