package userdomain

import "errors"

var (
	ErrInvalidEmail       = errors.New("Invalid email format")
	ErrInvalidPassword    = errors.New("Invalid password format. Minumin 3 characters")
	ErrEmailIsRequired    = errors.New("Email address is required")
	ErrFullNameIsRequired = errors.New("Full name is required")
	ErrPasswordIsRequired = errors.New("Password is required")
)
