package handlers

import (
	"profile-service/internal/repository"
	"profile-service/internal/service"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Handler struct {
	Srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{Srv: srv}
}

func (h *Handler)RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("POST /api/profile", h.CreateProfile)
}


func (h *Handler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var profReq repository.CreateProfileRequest
	err := json.NewDecoder(r.Body).Decode(&profReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prof, err := h.Srv.CreateProfile(r.Context(), msg)
	if err != nil {
		http.Error(w, "failed to create message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prof)
}

func (h *Handler) GetMessage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	msg, err := h.Srv.GetMessage(r.Context(), id)

	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to get message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func (h *Handler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var msg repository.Message
	if decodeErr := json.NewDecoder(r.Body).Decode(&msg); decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	err := h.Srv.UpdateMessage(r.Context(), id, msg.Text)

	if err != nil {
		if err.Error() == "ID not found" {
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to update message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.Srv.DeleteMessage(r.Context(), id)

	if err != nil {
		if err.Error() == "ID not found" {
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to delete message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
