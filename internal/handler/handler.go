package handler

import (
	"golangproj/internal/handler/authhandler"
	"golangproj/internal/handler/userhandler"
	"golangproj/internal/store/postgres"
	"golangproj/internal/usecase/auth"
	"golangproj/internal/usecase/user"
)

type Handlers struct {
	User userhandler.UserHandler
	Auth authhandler.AuthHandler
}

func NewHandlers(storage *postgres.PostgresStorage) *Handlers {
	return &Handlers{
		User: userhandler.UserHandler{
			Usecases: user.NewUserUsecases(storage.Users),
		},
		Auth: authhandler.AuthHandler{
			Usecases: auth.NewAuthUsecases(storage.Users),
		},
	}
}
