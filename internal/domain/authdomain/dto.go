// Package authdomain defines domain models and validation logic for authentication-related operations.
package authdomain

import "errors"

// LoginInput represents the user input required for authentication.
type LoginInput struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

// Validate checks that the required fields in LoginInput are present.
func (l *LoginInput) Validate() error {
	if l.EmailAddress == "" {
		return ErrEmailAddressRequired
	}

	if l.Password == "" {
		return ErrPasswordIsRequired
	}

	return nil
}

var (
	// ErrEmailAddressRequired is returned when the email field is empty.
	ErrEmailAddressRequired = errors.New("email is required")

	// ErrPasswordIsRequired is returned when the password field is empty.
	ErrPasswordIsRequired = errors.New("password is required")
)
