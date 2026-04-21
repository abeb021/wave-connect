package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(secret string, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//skip options
		if r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		var tokenString string

		// check for header
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			separated := strings.Split(authHeader, " ")
			if len(separated) == 2 && separated[0] == "Bearer" {
				tokenString = separated[1]
			}
		}

		// fallback to cookie (e.g. websocket)
		if tokenString == "" {
			cookie, err := r.Cookie("jwt")
			if err == nil {
				separated := strings.Split(cookie.Value, " ")
				if len(separated) == 2 && separated[0] == "Bearer" {
					tokenString = separated[1]
				}
			}
		}

		//no token
		if tokenString == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || token == nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			tokenValid := false
			if token != nil {
				tokenValid = token.Valid
			}
			log.Printf("JWT validation failed: %v, token.Valid: %v %v", err, tokenValid, tokenString)
			return
		}

		userID, err := token.Claims.GetSubject()
		if err != nil || userID == "" {
			http.Error(w, "Empty Token", http.StatusUnauthorized)
			return
		}

		// set user for later use
		r.Header.Set("X-User-ID", userID)
		ctx := context.WithValue(r.Context(), "userID", userID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
