package repository

import (
	"time"
)

type Profile struct {
	Username    string    `json:"username"`
	ID      string    `json:"id"`
	TimeCreated time.Time `json:"time_created"`
}

type CreateProfileRequest struct {
	Username string `json:"username"`
}

type ProfileRequest struct {
	Username string `json:"username"`
	ID   string `json:"id"`
}
