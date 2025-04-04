package auth

import (
	"context"
	"errors"
	"golangproj/internal/domain/userdomain"
	"golangproj/internal/usecase/user"
)

type RegisterUsecase interface {
	Execute(context.Context, *userdomain.UserInput) error
}

type Register struct {
	createUser user.CreateUserUsecase
}

func (r *Register) Execute(ctx context.Context, input *userdomain.UserInput) error {
	_, err := r.createUser.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

var (
	ErrInputIsNil = errors.New("Input is null.")
)
