package domain

import (
	"time"
)

type Profile struct {
	Username    string    `json:"username,omitempty"`
	ID          string    `json:"id"`
	TimeCreated time.Time `json:"time_created"`
	Bio         string    `json:"bio,omitempty"`
	Avatar      []byte    `json:"avatar,omitempty"`
}

type CreateProfileRequest struct {
	Username string `json:"username,omitempty"`
}

type ProfileRequest struct {
	Username string `json:"username,omitempty"`
	ID       string `json:"id"`
	Bio      string `json:"bio,omitempty"`
}
