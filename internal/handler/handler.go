// Package handler aggregates all HTTP handlers used throughout the application,
// initializing them with their respective use cases.
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

// Handlers holds all grouped HTTP handlers for different domains.
type Handlers struct {
	User userhandler.UserHandler
	Auth authhandler.AuthHandler
	Room roomhandler.RoomHandler
}

// NewHandlers initializes and returns all HTTP handlers by injecting their respective use cases.
func NewHandlers(storage *postgres.Storage, mongoStorage *mongodb.MongoStorage) *Handlers {
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
