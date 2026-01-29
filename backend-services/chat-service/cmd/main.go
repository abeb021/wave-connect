package main

import (
	"chat_service/api/handlers"
	"chat_service/api/middleware"
	"chat_service/internal/database"
	"context"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	dbPool, err := database.NewDB(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbPool.Close()

	if _, err = dbPool.Exec(ctx,
		`CREATE TABLE IF NOT EXISTS messages (
    	 	id SERIAL PRIMARY KEY,
     		text TEXT NOT NULL,
     		sender TEXT NOT NULL,
     		receiver TEXT NOT NULL,
     		time_sent TIMESTAMPTZ NOT NULL DEFAULT NOW()
	 	);`); err != nil {
		log.Fatal(err.Error())
	}
	
	h := &handlers.Handlers{DB: dbPool}

	r := http.NewServeMux()

	r.HandleFunc("GET /", h.HeIsInRoot)
	r.HandleFunc("POST /api/message", h.CreateMessage)
	r.HandleFunc("GET /api/message/{id}", h.GetMessage)
	r.HandleFunc("PUT /api/message/{id}", h.UpdateMessage)
	r.HandleFunc("DELETE /api/message/{id}", h.DeleteMessage)

	GlobalMiddleware := middleware.Chain(
		middleware.LoggingMiddleware,
		middleware.CorsMiddleware,
	)

	handler := GlobalMiddleware(r)

	server := http.Server {
		Addr: ":8080",
		Handler: handler,
	}

	log.Fatal(server.ListenAndServe())
}


