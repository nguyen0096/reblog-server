package middleware

import (
	"net/http"
	"reblog-server/service/user"
	"strings"

	"github.com/gorilla/mux"
)

func Authenticator(userService user.UserService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from request
			authorizationHeader := r.Header.Get("Authorization")
			if authorizationHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) != 2 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			isAuthenticated := userService.VerifyToken(bearerToken[1])
			if !isAuthenticated {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
