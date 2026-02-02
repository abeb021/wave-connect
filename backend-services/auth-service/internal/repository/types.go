package repository

import (
	"time"
)

type User struct {
	ID			string		 `json:"id"`
	Name		string		 `json:"text"`
    Email       string       `json:"email"`
	Password	string		 `json:"password"`
	TimeCreated time.Time	 `json:"timeSent"`
}
