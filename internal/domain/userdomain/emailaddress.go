package userdomain

import (
	"errors"
	"net/mail"
)

type EmailAddress struct {
	Value string
}

func NewEmailAddress(email string) (*EmailAddress, error) {
	parse, err := mail.ParseAddress(email)
	if err != nil {
		return nil, ErrInvalidEmail
	}

	return &EmailAddress{Value: parse.Address}, nil
}

var (
	ErrInvalidEmail = errors.New("Email is invalid")
)
