package user

import (
	"context"
	"errors"
	"morae/internal/dto/userdto"
	"morae/internal/mapper/usermapper"
	"morae/internal/store/postgres"
)

// ErrUserNotFound is returned when a user could not be found.
var ErrUserNotFound = errors.New("user not found")

// GetUserByIDUsecase defines the behavior for retrieving a user by ID.
type GetUserByIDUsecase interface {
	Execute(context.Context, uint64) (userdto.UserResponse, error)
}

// GetUserByID implements the GetUserByIDUsecase interface.
type GetUserByID struct {
	repo postgres.UserRepository
}

// Execute retrieves a user by ID and returns a UserResponse or an error.
func (u *GetUserByID) Execute(ctx context.Context, id uint64) (userdto.UserResponse, error) {
	entity, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return userdto.UserResponse{}, err
	}
	if entity == nil {
		return userdto.UserResponse{}, ErrUserNotFound
	}

	return usermapper.ToResponse(entity), nil
}
