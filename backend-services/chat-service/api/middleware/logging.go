package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := &responseWriter{ResponseWriter: w}
		
		next.ServeHTTP(ww, r)

		log.Printf(
			"%s %s %d %s %s",
			r.Method,
			r.URL.Path,
			ww.status,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(statusCode int){
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}