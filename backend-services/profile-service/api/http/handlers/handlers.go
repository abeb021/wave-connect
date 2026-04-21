package handlers

import (
	"encoding/json"
	"net/http"
	"profile-service/internal/repository"
	"profile-service/internal/service"
	"profile-service/usecases"

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

	id := r.Header.Get("X-User-ID")
	if id == "" {
		http.Error(w, "missing user id", http.StatusUnauthorized)
		return
	}
	prof, err := h.Srv.CreateProfile(r.Context(), &profReq, id)
	if err != nil {
		if err == usecases.ErrUserTaken {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "failed to create profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prof)
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	prof, err := h.Srv.GetProfile(r.Context(), id)

	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to get profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prof)
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("X-User-ID")

	var prof repository.Profile
	if decodeErr := json.NewDecoder(r.Body).Decode(&prof); decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	prof.ID = userId

	err := h.Srv.UpdateProfile(r.Context(), &prof)

	if err != nil {
		if err.Error() == "ID/Username not found" {
			http.Error(w, "ID/Username not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to update profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (h *Handler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")

	err := h.Srv.DeleteProfile(r.Context(), userID)

	if err != nil {
		if err.Error() == "ID not found" {
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to delete profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
