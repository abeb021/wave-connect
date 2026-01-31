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
	r.HandleFunc("POST /api/user", h.CreateUser)
	handler := http.NewServeMux()

	server := http.Server{
		Addr: ":8081",
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}