package handlers

import (
	"chat-service/internal/repository"
	"chat-service/internal/service"
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

func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msgReq repository.MessageRequest
	err := json.NewDecoder(r.Body).Decode(&msgReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msgReq.Sender = r.Header.Get("X-User-ID")

	msg, err := h.Srv.CreateMessage(r.Context(), &msgReq)
	if err != nil {
		http.Error(w, "failed to create message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}



func (h *Handler) GetConversation(w http.ResponseWriter, r *http.Request) {
	senderID := r.Header.Get("X-User-ID")

	if senderID == ""{
        http.Error(w, "your id is required", http.StatusBadRequest)
        return
    }

	conv, err := h.Srv.GetConversation(r.Context(), senderID)
	if err != nil{
		http.Error(w, "failed to get conversation", http.StatusInternalServerError)
		return
	}

	if conv == nil{
		conv = []repository.Message{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conv)
}

func (h *Handler) GetConversationWithPeer(w http.ResponseWriter, r *http.Request) {
	receiverID := r.PathValue("peerID")
	senderID := r.Header.Get("X-User-ID")

	if receiverID == ""{
        http.Error(w, "peer id is required", http.StatusBadRequest)
        return
    }

	conv, err := h.Srv.GetConversationWithPeer(r.Context(), senderID, receiverID)
	if err != nil{
		http.Error(w, "failed to get conversation", http.StatusInternalServerError)
		return
	}

	if conv == nil{
		conv = []repository.Message{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conv)
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

	msg.Sender = r.Header.Get("X-User-ID")

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
	senderID := r.Header.Get("X-User-ID")

	err := h.Srv.DeleteMessage(r.Context(), id, senderID)

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
