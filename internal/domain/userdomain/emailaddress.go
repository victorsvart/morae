// Package userdomain contains domain models and validation logic for users.
package userdomain

import (
	"errors"
	"net/mail"
)

// EmailAddress represents a validated email address.
type EmailAddress struct {
	Value string
}

// NewEmailAddress validates the given string and returns an EmailAddress instance if valid.
func NewEmailAddress(email string) (*EmailAddress, error) {
	parsed, err := mail.ParseAddress(email)
	if err != nil {
		return nil, ErrInvalidEmail
	}

	return &EmailAddress{Value: parsed.Address}, nil
}

// ErrInvalidEmail is returned when the provided email is not a valid format.
var ErrInvalidEmail = errors.New("email is invalid")
