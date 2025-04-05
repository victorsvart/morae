// Package room contains use cases related to room operations.
package room

import (
	"context"
	"errors"
	"morae/internal/domain/roomdomain"
	"morae/internal/dto/roomdto"
	"morae/internal/mapper/roommapper"
	"morae/internal/store/mongodb"
)

// UpdateRoomUsecase defines the contract for updating a room.
type UpdateRoomUsecase interface {
	Execute(context.Context, *roomdto.RoomDto) (roomdomain.Room, error)
}

// UpdateRoom implements the UpdateRoomUsecase using a MongoDB repository.
type UpdateRoom struct {
	repo mongodb.RoomRepository
}

// Execute updates an existing room in the database using the provided DTO.
func (u *UpdateRoom) Execute(ctx context.Context, input *roomdto.RoomDto) (roomdomain.Room, error) {
	document, err := roommapper.FromDto(input)
	if err != nil {
		return roomdomain.Room{}, err
	}

	err = u.repo.UpdateRoom(ctx, &document)
	if err != nil {
		return roomdomain.Room{}, errors.New("failed to update room")
	}

	return roommapper.ToDomain(&document), nil
}
