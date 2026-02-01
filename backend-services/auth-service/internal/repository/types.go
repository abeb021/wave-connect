package repository

import (
	"net/mail"
	"time"
)

type User struct {
	ID			string		 `json:"id"`
	Name		string		 `json:"text"`
    Email       mail.Address `json:"email"`
	Password	string		 `json:"password"`
	TimeCreated time.Time	 `json:"timeSent"`
}
