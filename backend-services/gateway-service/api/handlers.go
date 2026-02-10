package handlers

import (
	"gateway-service/api/middleware"
	"gateway-service/internal/service"
	"net/http"
)

func RegisterRoutes(r *http.ServeMux, srv *service.Service) {
	r.Handle("POST /api/auth/register", srv.AuthProxy())
	r.Handle("POST /api/auth/login", srv.AuthProxy())

	r.Handle("/api/auth", middleware.JWTMiddleware(srv.JWTSecret, srv.AuthURL, srv.AuthProxy()))
	r.Handle("/api/message", middleware.JWTMiddleware(srv.JWTSecret, srv.AuthURL, srv.ChatProxy()))
}
