// Package userdto defines user-related Data Transfer Objects and validation errors.
package userdto

import "errors"

var (
	// ErrEmailIsRequired indicates that the email address field is missing.
	ErrEmailIsRequired = errors.New("email address is required")

	// ErrFullNameIsRequired indicates that the full name field is missing.
	ErrFullNameIsRequired = errors.New("full name is required")

	// ErrPasswordIsRequired indicates that the password field is missing.
	ErrPasswordIsRequired = errors.New("password is required")
)
