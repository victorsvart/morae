package userdomain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Value string
}

func SetupPassword(plainPassword string) (*Password, error) {
	if len(plainPassword) < 3 {
		return nil, ErrInvalidPassword
	}

	password := &Password{Value: plainPassword}
	err := password.HashPassword()
	if err != nil {
		return nil, err
	}

	return password, nil
}

func (p *Password) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p.Value), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Value = string(hashed)
	return nil
}

func (p *Password) ComparePassword(plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.Value), []byte(plainPassword))
	if err != nil {
		return err
	}

	return nil
}

var (
	ErrInvalidPassword = errors.New("Invalid password. Minimum is 3 characters")
)
