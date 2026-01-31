package main

import (
	"auth-service/api/handlers"
	"log"
	"net/http"
)


func main (){


	h := &handlers.Handler{}

	r := http.NewServeMux()
	r.HandleFunc("POST /api/user", h.CreateUser)
	r.HandleFunc("GET /api/user/{id}", h.CreateUser)
	r.HandleFunc("PUT /api/user{id}", h.CreateUser)
	r.HandleFunc("DELETE /api/user{id}", h.CreateUser)

	handler := http.NewServeMux()

	server := http.Server{
		Addr: ":8081",
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}