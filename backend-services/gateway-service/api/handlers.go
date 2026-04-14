package handlers

import (
	"gateway-service/api/middleware"
	"gateway-service/internal/service"
	"net/http"
)

func RegisterRoutes(r *http.ServeMux, srv *service.Service) {
	r.Handle("POST /api/auth/register", srv.AuthProxy())
	r.Handle("POST /api/auth/login", srv.AuthProxy())

	r.Handle("/api/auth/", middleware.JWTMiddleware(srv.JWTSecret, srv.AuthProxy()))
	r.Handle("/api/feed/", middleware.JWTMiddleware(srv.JWTSecret, srv.FeedProxy()))
	r.Handle("/api/profile/", middleware.JWTMiddleware(srv.JWTSecret, srv.ProfileProxy()))
	r.Handle("/api/chat/", middleware.JWTMiddleware(srv.JWTSecret, srv.ChatProxy()))
}
