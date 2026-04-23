package domain

import "errors"

var (
	ErrProfileNotFound  = errors.New("profile not found")
	ErrUsernameTaken     = errors.New("username already taken")
)
