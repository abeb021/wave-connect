package repository

import (
	"time"
)

type Publication struct {
	PostID   string    `json:"post_id"`
	Text     string    `json:"text"`
	UserID   string    `json:"user_id"`
	TimeSent time.Time `json:"timeSent"`
}

type PublicationRequest struct {
	Text     string    `json:"text"`
}
