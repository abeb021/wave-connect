package main

import (
	"chat-service/api/handlers"
	"chat-service/api/middleware"
	"chat-service/internal/config"
	"chat-service/internal/repository"
	"chat-service/internal/service"
	"context"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()

	m, err := migrate.New("file://migrations", cfg.DatabaseURL())
	if err != nil {
		log.Fatalf("migrate setup: %v", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migrate up: %v", err)
	}

	dbPool, err := repository.NewPool(ctx, cfg)
	if err != nil {
		log.Fatalf("db pool: %v", err)
	}
	defer dbPool.Close()

	repo := repository.NewRepository(dbPool)
	srv := service.NewService(repo)
	h := &handlers.Handler{Srv: srv}

	r := http.NewServeMux()

	r.HandleFunc("POST /api/message", h.CreateMessage)
	r.HandleFunc("GET /api/message/{id}", h.GetMessage)
	r.HandleFunc("PUT /api/message/{id}", h.UpdateMessage)
	r.HandleFunc("DELETE /api/message/{id}", h.DeleteMessage)

	GlobalMiddleware := middleware.Chain(
		middleware.LoggingMiddleware,
		middleware.CorsMiddleware,
	)

	handler := GlobalMiddleware(r)

	server := http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: handler,
	}

	log.Fatal(server.ListenAndServe())
}
