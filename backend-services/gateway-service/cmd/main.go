package main

import (
	"gateway-service/api/middleware"
	"gateway-service/internal/config"
	"gateway-service/internal/proxy"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	cfg := config.Load()
	proxy := proxy.NewProxy(cfg.AuthServiceURL, cfg.ChatServiceURL)

	handler := middleware.Chain(
		middleware.CorsMiddleware,
		middleware.LoggingMiddleware,
	) (proxy)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func (){
		log.Fatal(server.ListenAndServe())
	}()
	
	down := make(chan os.Signal, 1)
	signal.Notify(down, syscall.SIGTERM, syscall.SIGINT)
	<-down
}