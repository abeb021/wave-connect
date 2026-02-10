package repository

import (
	"time"
)

type Message struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	TimeSent time.Time `json:"timeSent"`
}
