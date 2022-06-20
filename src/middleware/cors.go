package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func CorsMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		corsMiddleware := cors.New(cors.Options{
			AllowedOrigins: []string{"https://www.google.com"},
			AllowedMethods: []string{"GET", "POST", "PUT", "Delete"},
			AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
			Debug:          true,
		})
		corsMiddleware.Handler(next)
		next.ServeHTTP(w, r)
	}

}
