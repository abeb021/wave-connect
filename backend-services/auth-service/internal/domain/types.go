package domain

import (
	"time"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserDB struct {
	ID           string    `db:"id"`
	Email        string    `db:"email"`
	PasswordHASH string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}
