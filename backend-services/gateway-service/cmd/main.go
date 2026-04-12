package main

import (
	"gateway-service/api"
	"gateway-service/api/middleware"
	"gateway-service/internal/config"
	"gateway-service/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.Load()

	service := service.NewService(
		cfg.AuthServiceURL,
		cfg.ChatServiceURL,
		cfg.FeedServiceURL,
		cfg.ProfileServiceURL,
		cfg.JWTSecret,
	)

	handler := http.NewServeMux()

	handlers.RegisterRoutes(handler, service)

	middleware.Chain(
		middleware.CorsMiddleware,
		middleware.LoggingMiddleware,
	)(handler)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	down := make(chan os.Signal, 1)
	signal.Notify(down, syscall.SIGTERM, syscall.SIGINT)
	<-down
}
