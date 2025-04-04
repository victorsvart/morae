package handler

import (
	"morae/internal/handler/authhandler"
	"morae/internal/handler/userhandler"
	"morae/internal/store/postgres"
	"morae/internal/usecase/auth"
	"morae/internal/usecase/user"
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
