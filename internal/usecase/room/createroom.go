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

// CreateRoomInterface defines the contract for creating a new room.
type CreateRoomInterface interface {
	Execute(context.Context, *roomdto.RoomInput) (roomdomain.Room, error)
}

// CreateRoom implements CreateRoomInterface using a MongoDB repository.
type CreateRoom struct {
	repo mongodb.RoomRepository
}

// Execute creates a new room in the database from the provided input.
func (c *CreateRoom) Execute(ctx context.Context, input *roomdto.RoomInput) (roomdomain.Room, error) {
	document := roommapper.FromInput(input)

	err := c.repo.CreateRoom(ctx, &document)
	if err != nil {
		return roomdomain.Room{}, errors.New("failed to create room")
	}

	domain := roommapper.ToDomain(&document)
	return domain, nil
}
