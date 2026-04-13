package handlers

import (
	"feed-service/internal/repository"
	"feed-service/internal/service"
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

func (h *Handler) CreatePublication(w http.ResponseWriter, r *http.Request) {
	var pubReq repository.PublicationRequest
	err := json.NewDecoder(r.Body).Decode(&pubReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pubReq.UserID = r.Header.Get("X-User-ID")

	pub, err := h.Srv.CreatePublication(r.Context(), pubReq)

	if err != nil {
		http.Error(w, "failed to create publication", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pub)
}

func (h *Handler) GetPublication(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pub, err := h.Srv.GetPublication(r.Context(), id)

	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to get publication", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pub)
}

func (h *Handler) UpdatePublication(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var pub repository.PublicationRequest
	if decodeErr := json.NewDecoder(r.Body).Decode(&pub); decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	err := h.Srv.UpdatePublication(r.Context(), id, pub.Text)

	if err != nil {
		if err.Error() == "ID not found" {
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to update publication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (h *Handler) DeletePublication(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.Srv.DeletePublication(r.Context(), id)

	if err != nil {
		if err.Error() == "ID not found" {
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to delete publication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
