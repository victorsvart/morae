package userdomain

import "errors"

var (
	ErrInvalidEmail    = errors.New("Invalid email format")
	ErrInvalidPassword = errors.New("Invalid password format. Minumin 3 characters")
)
