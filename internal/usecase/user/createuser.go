package user

import (
	"context"
	"errors"
	"morae/internal/domain/userdomain"
	"morae/internal/mapper/usermapper"
	"morae/internal/store/postgres"
)

type CreateUserUsecase interface {
	Execute(context.Context, *userdomain.UserInput) (*userdomain.UserResponse, error)
}

type Create struct {
	repo postgres.UserRepository
}

func (c *Create) Execute(ctx context.Context, input *userdomain.UserInput) (*userdomain.UserResponse, error) {
	domain, err := usermapper.FromInput(input)
	if err != nil {
		return nil, err
	}

	entity := usermapper.ToEntity(domain)
	if err := c.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return usermapper.ToResponse(entity), nil
}

var (
	ErrInputIsNil = errors.New("Input is null.")
)
