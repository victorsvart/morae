package auth

import (
	"morae/internal/store/postgres"
	"morae/internal/usecase/user"
)

type AuthUsecases struct {
	Login    LoginUsecase
	Register RegisterUsecase
}

func NewAuthUsecases(repo postgres.UserRepository) *AuthUsecases {
	return &AuthUsecases{
		Login:    &Login{repo},
		Register: &Register{user.NewUserUsecases(repo).Create}, // reusing the CreateUsecase
	}
}
