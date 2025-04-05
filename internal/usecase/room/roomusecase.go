// Package room contains business logic for handling room-related operations.
package room

import "morae/internal/store/mongodb"

// Usecases holds all use case implementations for room operations.
type Usecases struct {
	GetAllRooms GetAllRoomsUsecase
	GetByID     GetRoomByIDUsecase
	CreateRoom  CreateRoomInterface
	UpdateRoom  UpdateRoomUsecase
	DeleteRoom  DeleteRoomUsecase
}

// NewRoomUsecases initializes and returns a new instance of RoomUsecases.
func NewRoomUsecases(repo mongodb.RoomRepository) *Usecases {
	return &Usecases{
		GetAllRooms: &GetAllRooms{repo},
		GetByID:     &GetRoomByID{repo},
		CreateRoom:  &CreateRoom{repo},
		UpdateRoom:  &UpdateRoom{repo},
		DeleteRoom:  &DeleteRoom{repo},
	}
}
