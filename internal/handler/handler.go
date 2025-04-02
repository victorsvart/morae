package handler

import (
	"golangproj/internal/handler/userhandler"
	"golangproj/internal/store/postgres"
	"golangproj/internal/usecase/user"
)

type Handlers struct {
	User userhandler.UserHandler
}

func NewHandlers(storage *postgres.PostgresStorage) *Handlers {
	return &Handlers{
		User: userhandler.UserHandler{
			Usecases: user.NewUserUsecases(storage.Users),
		},
	}
}
