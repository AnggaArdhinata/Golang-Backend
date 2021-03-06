package middleware

import (
	"backend-golang/src/helpers"
	"net/http"
	"strings"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			helpers.New("Invalid header type", 401, false).Send(w)
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)
		checkToken, err := helpers.CheckToken(token)
		if err != nil {
			helpers.New(err.Error(), 201, true).Send(w)
			return
		}
		if !checkToken {
			helpers.New("Silahkan login kembali", 401, false).Send(w)
			return
		}
		next.ServeHTTP(w, r)
	}
}