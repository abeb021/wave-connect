package repository

import (
	"time"
)

type MessageRequest struct {
	Text     string    `json:"text"`
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
}

type Message struct{
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	TimeSent time.Time `json:"timeSent"`
}