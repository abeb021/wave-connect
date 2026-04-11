package repository

import (
	"time"
)

type Profile struct {
	Username    string    `json:"username"`
	UserID      string    `json:"user_id"`
	TimeCreated time.Time `json:"timeCreated"`
}

type CreateProfileRequest struct {
	Username string `json:"text"`
}
