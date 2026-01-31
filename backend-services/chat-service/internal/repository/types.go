package repository

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	TimeSent time.Time `json:"timeSent"`
}