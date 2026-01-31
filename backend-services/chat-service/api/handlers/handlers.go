package handlers

import (
	"chat-service/internal/repository"
	"chat-service/internal/service"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Handler struct{
	Srv *service.Service
}

func (h *Handler)CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg repository.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg, err = h.Srv.CreateMessage(r.Context(), msg)
	if err != nil{
		http.Error(w, "failed to create message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func (h *Handler)GetMessage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	
	msg, err := h.Srv.GetMessage(r.Context(), uuid.MustParse(id))
	
	if err != nil{
		if err == pgx.ErrNoRows{
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to get message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func (h *Handler)UpdateMessage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var msg repository.Message
	if 	decodeErr := json.NewDecoder(r.Body).Decode(&msg); decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	err := h.Srv.UpdateMessage(r.Context(), uuid.MustParse(id), msg.Text)

	if err != nil {
		if err == errors.New("ID not found") {
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to update message", http.StatusInternalServerError)
		return
	}


	w.WriteHeader(http.StatusNoContent)

}

func (h *Handler)DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	
	err := h.Srv.DeleteMessage(r.Context(), uuid.MustParse(id))

	if err != nil {
		if err == errors.New("ID not found") {
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to delete message", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}
