// Package auth provides authentication-related use cases such as user registration.
package auth

import (
	"context"
	"errors"
	"morae/internal/dto/userdto"
	"morae/internal/usecase/user"
)

// RegisterUsecase defines the interface for user registration logic.
type RegisterUsecase interface {
	Execute(context.Context, *userdto.UserInput) error
}

// Register implements the RegisterUsecase using the user.CreateUserUsecase.
type Register struct {
	createUser user.CreateUserUsecase
}

// Execute handles the user registration by delegating to the CreateUserUsecase.
func (r *Register) Execute(ctx context.Context, input *userdto.UserInput) error {
	if input == nil {
		return ErrInputIsNil
	}

	_, err := r.createUser.Execute(ctx, input)
	return err
}

// ErrInputIsNil is returned when the registration input is nil.
var ErrInputIsNil = errors.New("input is nil")
