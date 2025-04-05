package user

import (
	"context"
	"morae/internal/domain/userdomain"
	"morae/internal/dto/userdto"
	"morae/internal/mapper/usermapper"
	"morae/internal/store/postgres"
)

// UpdateUserUsecase defines the interface for updating a user.
type UpdateUserUsecase interface {
	Execute(context.Context, *userdomain.User) (userdto.UserResponse, error)
}

// Update implements the UpdateUserUsecase interface using a Postgres repository.
type Update struct {
	repo postgres.UserRepository
}

// Execute updates a user in the repository and returns the updated user response.
func (u *Update) Execute(ctx context.Context, userDomain *userdomain.User) (userdto.UserResponse, error) {
	entity := usermapper.ToEntity(userDomain)
	if err := u.repo.Update(ctx, &entity); err != nil {
		return userdto.UserResponse{}, err
	}

	return usermapper.ToResponse(&entity), nil
}
