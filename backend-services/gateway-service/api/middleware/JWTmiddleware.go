package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(secret, authServiceURL string, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//skip options
		if r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		separated := strings.Split(authHeader, " ")
		if len(separated) != 2 || separated[0] != "Bearer" {
			http.Error(w, "Malformed authorization header", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(separated[1], &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			log.Printf("JWT validation failed: %v, token.Valid: %v", err, token.Valid)
			return
		}

		userID, err := token.Claims.GetSubject()
		if err != nil || userID == "" {
			http.Error(w, "Empty Token", http.StatusUnauthorized)
			return
		}

		authURL := fmt.Sprintf("%s/api/auth/%s", authServiceURL, userID)
		req, err := http.NewRequest(http.MethodGet, authURL, nil)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to identify user", http.StatusUnauthorized)
			return
		}
		resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			http.Error(w, "Could not find the user", http.StatusUnauthorized)
			return
		}
		if resp.StatusCode != http.StatusAccepted {
			http.Error(w, "Failed to identify user", http.StatusUnauthorized)
			return
		}
		r.Header.Set("X-User-ID", userID)
		ctx := context.WithValue(r.Context(), "userID", userID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
