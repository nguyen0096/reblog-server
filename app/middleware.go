package app

import (
	"context"
	"log"
	"net/http"
)

func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		ctx := context.WithValue(r.Context(), "username", "Nguyen")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
