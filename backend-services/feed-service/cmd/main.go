package main

import (
	"feed-service/api/handlers"
	"feed-service/api/middleware"
	"feed-service/internal/config"
	"feed-service/internal/repository"
	"feed-service/internal/service"
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

	r.HandleFunc("POST /api/feed/", h.CreatePublication)
	r.HandleFunc("GET /api/feed/{id}", h.GetPublication)
	r.HandleFunc("PUT /api/feed/{id}", h.UpdatePublication)
	r.HandleFunc("DELETE /api/feed/{id}", h.DeletePublication)

	GlobalMiddleware := middleware.Chain(
		middleware.LoggingMiddleware,
		middleware.CorsMiddleware,
	)

	handler := GlobalMiddleware(r)

	server := http.Server{
		Addr:    ":8083",
		Handler: handler,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	down := make(chan os.Signal, 1)
	signal.Notify(down, syscall.SIGTERM, syscall.SIGINT)
	<-down
}
