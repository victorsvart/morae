// Package user provides use cases for managing user entities, including user creation and interaction with the persistence layer.
package user

import (
	"context"
	"errors"
	"morae/internal/dto/userdto"
	"morae/internal/mapper/usermapper"
	"morae/internal/store/postgres"
)

// CreateUserUsecase defines the interface for creating a new user.
type CreateUserUsecase interface {
	Execute(context.Context, *userdto.UserInput) (userdto.UserResponse, error)
}

// Create implements the CreateUserUsecase using a Postgres repository.
type Create struct {
	repo postgres.UserRepository
}

// Execute handles the creation of a new user, converting input data to domain and entity layers.
func (c *Create) Execute(ctx context.Context, input *userdto.UserInput) (userdto.UserResponse, error) {
	domain, err := usermapper.FromInput(input)
	if err != nil {
		return userdto.UserResponse{}, err
	}

	entity := usermapper.ToEntity(&domain)
	if err := c.repo.Create(ctx, &entity); err != nil {
		return userdto.UserResponse{}, err
	}

	return usermapper.ToResponse(&entity), nil
}

// ErrInputIsNil is returned when a nil input is passed to the usecase.
var (
	ErrInputIsNil = errors.New("input is null")
)
