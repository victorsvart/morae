package user

import (
	"context"
	"errors"
	"morae/internal/dto/userdto"
	"morae/internal/mapper/usermapper"
	"morae/internal/store/postgres"
)

type GetUserByIdUsecase interface {
	Execute(context.Context, uint64) (userdto.UserResponse, error)
}

type GetUserById struct {
	repo postgres.UserRepository
}

func (u *GetUserById) Execute(ctx context.Context, id uint64) (userdto.UserResponse, error) {
	entity, err := u.repo.GetById(ctx, id)

	if entity == nil {
		return userdto.UserResponse{}, errors.New("User not found")
	}

	if err != nil {
		return userdto.UserResponse{}, err
	}

	return usermapper.ToResponse(entity), nil
}

