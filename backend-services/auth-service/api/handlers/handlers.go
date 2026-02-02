package handlers

import (
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"encoding/json"
	"net/http"
	"net/mail"
	"time"

	"github.com/jackc/pgx/v5"
)

type Handler struct{
	Srv *service.Service
}

func NewHandler (srv *service.Service) *Handler{
    return &Handler{Srv: srv}
}

func (h *Handler)Register(w http.ResponseWriter, r *http.Request){
	var usr repository.User
    err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.Srv.Register(r.Context(), usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("token", token)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler)Login(w http.ResponseWriter, r *http.Request){	

}

func (h *Handler)DeleteUser(w http.ResponseWriter, r *http.Request){

}

func (h *Handler)GetUser(w http.ResponseWriter, r *http.Request){

}