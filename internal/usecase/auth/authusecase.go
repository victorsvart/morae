// Package auth provides use cases for authentication such as login and registration.
package auth

import (
	"morae/internal/store/postgres"
	"morae/internal/usecase/user"
)

// Usecases holds all authentication-related use cases.
type Usecases struct {
	Login    LoginUsecase
	Register RegisterUsecase
}

// NewAuthUsecases creates a new instance of AuthUsecases with required dependencies.
func NewAuthUsecases(repo postgres.UserRepository) *Usecases {
	return &Usecases{
		Login:    &Login{repo},
		Register: &Register{user.NewUserUsecases(repo).Create},
	}
}
