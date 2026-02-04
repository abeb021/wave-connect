package repository

import (
	"time"
)

type UserRequest struct {
	Username    string		 `json:"username"`
    Email       string       `json:"email"`
	Password	string		 `json:"password"`
}

type UserResponse struct {
	ID			string		 `json:"id"`
	Username	string		 `json:"username"`
    Email       string       `json:"email"`
	CreatedAt	time.Time	 `json:"created_at"`
}

type UserDB struct {
	ID				string		 `db:"id"`
	Username		string		 `db:"username"`
    Email       	string       `db:"email"`
	PasswordHASH	string		 `db:"password_hash"`
	CreatedAt 		time.Time	 `db:"created_at"`
}
