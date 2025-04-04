package auth

import (
	"context"
	"errors"
	"morae/internal/dto/userdto"
	"morae/internal/usecase/user"
)

type RegisterUsecase interface {
	Execute(context.Context, *userdto.UserInput) error
}

type Register struct {
	createUser user.CreateUserUsecase
}

func (r *Register) Execute(ctx context.Context, input *userdto.UserInput) error {
	_, err := r.createUser.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

var (
	ErrInputIsNil = errors.New("Input is null.")
)
