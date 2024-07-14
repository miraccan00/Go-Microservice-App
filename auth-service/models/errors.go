package models

import "errors"

var (
	ErrUsernameShort = errors.New("username is too short, should be at least 3 characters")
	ErrUsernameTaken = errors.New("username is already taken")
)
