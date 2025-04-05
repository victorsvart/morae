// Package room contains use cases related to room operations.
package room

import (
	"context"
	"morae/internal/store/mongodb"
)

// DeleteRoomUsecase defines the contract for deleting a room.
type DeleteRoomUsecase interface {
	Execute(context.Context, string) error
}

// DeleteRoom implements the DeleteRoomUsecase using a MongoDB repository.
type DeleteRoom struct {
	repo mongodb.RoomRepository
}

// Execute deletes a room by its ID.
func (d *DeleteRoom) Execute(ctx context.Context, id string) error {
	return d.repo.DeleteRoom(ctx, id)
}
