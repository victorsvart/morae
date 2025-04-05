package userdomain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Password represents a user's password with related behavior.
type Password struct {
	Value string
}

// SetupPassword creates a new Password and hashes it if valid.
func SetupPassword(plain string) (*Password, error) {
	if len(plain) < 3 {
		return nil, ErrInvalidPassword
	}

	p := &Password{Value: plain}
	if err := p.HashPassword(); err != nil {
		return nil, err
	}

	return p, nil
}

// HashPassword hashes the password's value using bcrypt.
func (p *Password) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p.Value), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Value = string(hashed)
	return nil
}

// ComparePassword checks if the provided plain password matches the hashed value.
func (p *Password) ComparePassword(plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.Value), []byte(plain))
}

var (
	// ErrInvalidPassword is returned when a password doesn't meet length requirements.
	ErrInvalidPassword = errors.New("invalid password: minimum length is 3 characters")
)
