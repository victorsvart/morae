package userdto

import "errors"

var (
	ErrEmailIsRequired    = errors.New("Email address is required")
	ErrFullNameIsRequired = errors.New("Full name is required")
	ErrPasswordIsRequired = errors.New("Password is required")
)
