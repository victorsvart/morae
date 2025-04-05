// Package auth provides authentication-related use cases such as user login.
package auth

import (
	"context"
	"errors"
	"morae/internal/domain/authdomain"
	"morae/internal/mapper/usermapper"
	"morae/internal/store/postgres"

	"golang.org/x/crypto/bcrypt"
)

// LoginUsecase defines the interface for user login logic.
type LoginUsecase interface {
	Execute(context.Context, *authdomain.LoginInput) error
}

// Login implements the LoginUsecase using the UserRepository.
type Login struct {
	repo postgres.UserRepository
}

// Execute handles user login by validating credentials against stored user data.
func (l *Login) Execute(ctx context.Context, input *authdomain.LoginInput) error {
	if input == nil {
		return ErrInputIsNil
	}

	userEntity, err := l.repo.FindByEmail(ctx, input.EmailAddress)
	if err != nil {
		return err
	}

	userDomain := usermapper.ToDomain(userEntity)
	if err := userDomain.Password.ComparePassword(input.Password); err != nil {
		if !errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return err
		}
		return ErrInvalidCredentials
	}

	return nil
}

var (
	// ErrInvalidCredentials is returned when the provided login credentials are incorrect.
	ErrInvalidCredentials = errors.New("invalid credentials")
)
