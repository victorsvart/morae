package user

import (
	"golangproj/internal/store/postgres"
)

type UserUsecases struct {
	Create  CreateUserUsecase
	GetById GetUserByIdUsecase
}

func NewUserUsecases(repo postgres.UserRepository) *UserUsecases {
	return &UserUsecases{
		Create:  &Create{repo},
		GetById: &GetUserById{repo},
	}
}
