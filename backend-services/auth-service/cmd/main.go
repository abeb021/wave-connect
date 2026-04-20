package main

import (
	"auth-service/api/handlers"
	"auth-service/api/middleware"
	"auth-service/internal/config"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/pkg/jwt"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	//context and cfg startup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := config.Load()

	//migrations
	m, err := migrate.New("file://migrations", cfg.DatabaseURL())
	if err != nil {
		log.Fatalf("migrate setup: %v", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migrate up: %v", err)
	}

	//db pool connection
	dbPool, err := repository.NewPool(ctx, cfg.DatabaseURL())
	if err != nil {
		log.Fatalf("db pool: %v", err)
	}
	defer dbPool.Close()

	//setup jwt service
	auth := jwt.NewAuthService(cfg.JWTSecret)
	//backend architecture setup
	repo := repository.NewRepository(dbPool)
	srv := service.NewService(repo, auth)
	h := handlers.NewHandler(srv)

	//routing
	r := http.NewServeMux()

	r.HandleFunc("POST /api/auth/register", h.Register)
	r.HandleFunc("POST /api/auth/login", h.Login)
	r.HandleFunc("GET /api/auth/{id}", h.GetUser)
	r.HandleFunc("DELETE /api/auth/{id}", h.DeleteUser)

	GlobalMiddleware := middleware.Chain(
		middleware.LoggingMiddleware,
	)

	handler := GlobalMiddleware(r)

	server := http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	down := make(chan os.Signal, 1)
	signal.Notify(down, syscall.SIGTERM, syscall.SIGINT)
	<-down
}
