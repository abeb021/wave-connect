package handlers

import (
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/usecases"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{Srv: srv}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var usr repository.UserRequest
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usrResponse, err := h.Srv.Register(r.Context(), &usr)
	if err != nil {
		if err == usecases.ErrUserTaken {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usrResponse)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var usr repository.UserRequest
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.Srv.Login(r.Context(), &usr)
	if err != nil {
		if err == usecases.ErrWrongPassword {
			http.Error(w, usecases.ErrWrongPassword.Error(), http.StatusBadRequest)
			return
		}
		if err == usecases.ErrUserNotFound {
			http.Error(w, usecases.ErrUserNotFound.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	

	w.Header().Set("Authorization", "Bearer " + token)
	w.WriteHeader(http.StatusAccepted)
}


func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	usrResponse, err := h.Srv.GetUserById(r.Context(), id)
	if err != nil {
		if err == usecases.ErrUserNotFound{
			http.Error(w, usecases.ErrUserNotFound.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(usrResponse)
}


func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.Srv.DeleteUser(r.Context(), id)
	if err != nil{
		if err == usecases.ErrUserNotFound {
			http.Error(w, usecases.ErrUserNotFound.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}


