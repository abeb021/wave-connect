package repository

import (
	"time"
)

type Publication struct {
	ID          string    `json:"id"`
	Text        string    `json:"text"`
	UserID      string    `json:"user_id"`
	TimeCreated time.Time `json:"time_created"`
}

type PublicationRequest struct {
	Text   string `json:"text"`
	UserID string `json:"user_id"`
}

type Comment struct {
	ID          string    `json:"id"`
	PubID       string    `json:"pub_id"`
	Text        string    `json:"text"`
	UserID      string    `json:"user_id"`
	TimeCreated time.Time `json:"time_created"`
}

type CommentRequest struct {
	Text   string `json:"text"`
	PubID  string `json:"pub_id"`
	UserID string `json:"user_id"`
}
