package middleware

import "net/http"

func GetOnly(mux http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				http.Error(w, "Invalid method", http.StatusBadRequest)
				return
			}
			mux.ServeHTTP(w, r)
		})
}
