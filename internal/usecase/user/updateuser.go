package user

import (
	"context"
	"golangproj/internal/domain/userdomain"
	"golangproj/internal/mapper/usermapper"
	"golangproj/internal/store/postgres"
)

type UpdateUserUsecase interface {
	Execute(context.Context, *userdomain.User) (*userdomain.UserResponse, error)
}

type Update struct {
	repo postgres.UserRepository
}

func (u *Update) Execute(ctx context.Context, userDomain *userdomain.User) (*userdomain.UserResponse, error) {
	entity := usermapper.ToEntity(userDomain)
	if err := u.repo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return usermapper.ToResponse(entity), nil
}
