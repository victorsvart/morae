package user

import (
	"context"
	"errors"
	"golangproj/internal/domain/userdomain"
	"golangproj/internal/mapper/usermapper"
	"golangproj/internal/store/postgres"
)

type GetUserByIdUsecase interface {
	Execute(context.Context, uint64) (*userdomain.UserResponse, error)
}

type GetUserById struct {
	repo postgres.UserRepository
}

func (u *GetUserById) Execute(ctx context.Context, id uint64) (*userdomain.UserResponse, error) {
	if id == 0 {
		return nil, ErrInvalidId
	}

	entity, err := u.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return usermapper.ToResponse(entity), nil
}

var (
	ErrInvalidId = errors.New("User id is invalid")
)
