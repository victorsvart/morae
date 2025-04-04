package authdomain

import "errors"

type LoginInput struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

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
	ErrEmailAddressRequired = errors.New("Email is required")
	ErrPasswordIsRequired   = errors.New("Password is required")
)
