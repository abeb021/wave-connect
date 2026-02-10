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

	//backend architecture setup
	repo := repository.NewRepository(dbPool)
	srv := service.NewService(repo)
	h := handlers.NewHandler(srv)

	//routing
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
		Addr:    ":8082",
		Handler: handler,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	down := make(chan os.Signal, 1)
	signal.Notify(down, syscall.SIGTERM, syscall.SIGINT)
	<-down
}
