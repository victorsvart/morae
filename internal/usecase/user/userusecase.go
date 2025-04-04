package user

import (
	"morae/internal/store/postgres"
)

type UserUsecases struct {
	GetById GetUserByIdUsecase
	Create  CreateUserUsecase
	Update  UpdateUserUsecase
	Delete  DeleteUserUsecase
}

func NewUserUsecases(repo postgres.UserRepository) *UserUsecases {
	return &UserUsecases{
		GetById: &GetUserById{repo},
		Create:  &Create{repo},
		Update:  &Update{repo},
		Delete:  &Delete{repo},
	}
}
