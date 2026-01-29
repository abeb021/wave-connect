package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Message struct {
	Id       int       `json:"id"`
	Text     string    `json:"text"`
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	TimeSent time.Time `json:"timeSent"`
}

type Handlers struct{
	DB *pgxpool.Pool
}

func (h *Handlers)HeIsInRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("omg are you here"))
}

func (h *Handlers)CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	row := h.DB.QueryRow(
		r.Context(), 
		`INSERT INTO messages (text, sender, receiver)
	 	 VALUES ($1, $2, $3)
	 	 RETURNING id, time_sent`, 
		msg.Text, msg.Sender, msg.Receiver)
	
	if err := row.Scan(&msg.Id, &msg.TimeSent); err != nil{
		http.Error(w, "failed to create message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func (h *Handlers)GetMessage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	var msg Message
	row := h.DB.QueryRow(
		r.Context(),
		`SELECT * 
		 FROM messages 
		 WHERE id = $1`, 
		id )
	err = row.Scan(&msg.Id, &msg.Text , &msg.Sender, &msg.Receiver, &msg.TimeSent )
	if err != nil{
		if err == pgx.ErrNoRows{
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		}

		http.Error(w, "failed to get message", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(msg.Text))
}

func (h *Handlers)UpdateMessage(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "failed to save message", http.StatusInternalServerError)
		return
	}
	if ct.RowsAffected() == 0 {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (h *Handlers)DeleteMessage(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "failed to save message", http.StatusInternalServerError)
		return
	}
	if ct.RowsAffected() == 0 {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}
