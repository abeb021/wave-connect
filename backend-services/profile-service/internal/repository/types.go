package repository

import (
	"time"
)

type Profile struct {
	Username    string    `json:"username"`
	UserID      string    `json:"user_id"`
	TimeCreated time.Time `json:"time_created"`
}

type CreateProfileRequest struct {
	Username string `json:"text"`
}

type ProfileRequest struct {
	Username string `json:"text"`
	UserID   string `json:"user_id"`
}
