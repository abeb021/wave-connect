package handlers

import (
	"auth-service/internal/service"
	"chat-service/internal/repository"
	"chat-service/internal/service"
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
    
}

func (h *Handler)Login(w http.ResponseWriter, r *http.Request){	

}

func (h *Handler)DeleteUser(w http.ResponseWriter, r *http.Request){

}

func (h *Handler)GetUser(w http.ResponseWriter, r *http.Request){

}