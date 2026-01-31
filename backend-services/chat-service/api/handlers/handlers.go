package handlers

import (
	"chat-service/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Message struct {
	ID       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	TimeSent time.Time `json:"timeSent"`
}

type Handler struct{
	Srv *service.Service
}

func (h *Handler)CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg.ID = uuid.New()
	row := h.DB.QueryRow(
		r.Context(), 
		`INSERT INTO messages (id, text, sender, receiver)
	 	 VALUES ($1, $2, $3, $4)
	 	 RETURNING time_sent`, 
		msg.ID, msg.Text, msg.Sender, msg.Receiver,
	)
	
	if err := row.Scan(&msg.TimeSent); err != nil{
		http.Error(w, "failed to create message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func (h *Handler)GetMessage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var msg Message
	row := h.DB.QueryRow(
		r.Context(),
		`SELECT id, text, sender, receiver, time_sent
		 FROM messages 
		 WHERE id = $1`, 
		id )
	err := row.Scan(&msg.ID, &msg.Text , &msg.Sender, &msg.Receiver, &msg.TimeSent )
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
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if 	err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var msg Message
	if 	decodeErr := json.NewDecoder(r.Body).Decode(&msg); decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	ct, err := h.DB.Exec(
		r.Context(),
		`UPDATE messages
		 SET text = $1
		 WHERE id = $2`,
		msg.Text, id,
	)

	if err != nil {
		http.Error(w, "failed to update message", http.StatusInternalServerError)
		return
	}
	if ct.RowsAffected() == 0 {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (h *Handler)DeleteMessage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	ct, err := h.DB.Exec(
		r.Context(),
		`DELETE FROM messages 
		 WHERE id=$1`,
		id,
	)
	if err != nil {
		http.Error(w, "failed to delete message", http.StatusInternalServerError)
		return
	}
	if ct.RowsAffected() == 0 {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}
