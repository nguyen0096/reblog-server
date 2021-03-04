package middleware

import (
	"net/http"
	"reblog-server/service/user"
	"reblog-server/utils"
	"strings"

	"github.com/gorilla/mux"
)

func Authenticator(userService user.UserService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method == "POST" && (r.URL.Path == "/user" || r.URL.Path == "/auth/login") {
				utils.Info("Path: %s", r.URL.Path)
				next.ServeHTTP(w, r)
				return
			}

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
