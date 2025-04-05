package handler

import (
	"morae/internal/handler/authhandler"
	"morae/internal/handler/roomhandler"
	"morae/internal/handler/userhandler"
	"morae/internal/store/mongodb"
	"morae/internal/store/postgres"
	"morae/internal/usecase/auth"
	"morae/internal/usecase/room"
	"morae/internal/usecase/user"
)

type Handlers struct {
	User userhandler.UserHandler
	Auth authhandler.AuthHandler
	Room roomhandler.RoomHandler
}

func NewHandlers(storage *postgres.PostgresStorage, mongoStorage *mongodb.MongoStorage) *Handlers {
	return &Handlers{
		User: userhandler.UserHandler{
			Usecases: user.NewUserUsecases(storage.Users),
		},
		Auth: authhandler.AuthHandler{
			Usecases: auth.NewAuthUsecases(storage.Users),
		},
		Room: roomhandler.RoomHandler{
			Usecases: room.NewRoomUsecases(mongoStorage.Rooms),
		},
	}
}
