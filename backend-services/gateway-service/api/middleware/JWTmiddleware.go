package middleware

import "net/http"

func JWTMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == ""{
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}