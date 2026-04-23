package domain

import "errors"

var (
	ErrMessageNotFound  = errors.New("user not found")
	ErrMessageTaken     = errors.New("username/email already taken")
)
