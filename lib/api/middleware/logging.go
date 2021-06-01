package middleware

import (
	"net/http"
	"reblog-server/utils"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Info(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Next()
// 	}
// }
