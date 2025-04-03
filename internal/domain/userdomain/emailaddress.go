package userdomain

import "net/mail"

type EmailAddress struct {
	Value string
}

func SetEmail(value string) *EmailAddress {
	return &EmailAddress{value}
}

func NewEmailAddress(email string) (*EmailAddress, error) {
	parse, err := mail.ParseAddress(email)
	if err != nil {
		return nil, ErrInvalidEmail
	}

	return &EmailAddress{Value: parse.Address}, nil
}
