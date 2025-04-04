package auth

import (
	"context"
	"errors"
	"morae/internal/domain/authdomain"
	"morae/internal/mapper/usermapper"
	"morae/internal/store/postgres"

	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase interface {
	Execute(context.Context, *authdomain.LoginInput) error
}

type Login struct {
	repo postgres.UserRepository
}

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
	ErrInvalidCredentials = errors.New("Invalid crendetials")
)
