package domain

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrUserTaken     = errors.New("username/email already taken")
	ErrWrongPassword = errors.New("wrong password")
)
