package handlers

import (
	"net/http"
	"net/mail"
	"time"

)


type User struct {
    ID           string       `json:"id"`
    Email        mail.Address    `json:"email"`
    Name         string          `json:"name"`
    RegisterTime time.Time       `json:"register_time"`
}

type Handler struct{
    
}

func (h *Handler)CreateUser(w http.ResponseWriter, r *http.Request){
    
}

func (h *Handler)GetUser(w http.ResponseWriter, r *http.Request){	
}

func (h *Handler)UpdateUser(w http.ResponseWriter, r *http.Request){
}
func (h *Handler)DeleteUser(w http.ResponseWriter, r *http.Request){
}